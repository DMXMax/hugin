package main

import (
	"flag"
	"fmt"
	 "math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"log"
	"github.com/DMXMax/noppa/dictionary"
	"github.com/DMXMax/noppa/weather"
	owm "github.com/briandowns/openweathermap"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	DiscordToken string
	WeatherToken string
	weatherOnly  *bool
)

type WindVelocity struct{
	KpH float64
	MpH float64
}

func init() {

	//	flag.StringVar(&Token, "t", "", "Bot Token")
	//	flag.Parse()
	weatherOnly = flag.Bool("weather", false, "Test Weather Only")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	//dg, err := discordgo.New("Bot " + Token)
	if WeatherToken = os.Getenv("NOPPA_WEATHER"); WeatherToken == "" {
		fmt.Println("No Weather Token")
		os.Exit(1)
	}

	if DiscordToken = os.Getenv("NOPPA_PIE"); DiscordToken == "" {
		fmt.Println("No token")
		os.Exit(1)
	}

	if *weatherOnly == true {
		log.Println("Testing Weather Only")
		w, err := owm.NewCurrent("C", "EN", WeatherToken)
		if err != nil {
			fmt.Println(err)
		} else {
			w.CurrentByID(4990729)
			log.Println(w)
			log.Printf("The current temperature in %s is %v degrees\n", w.Name, w.Main.Temp)
			log.Println(w.Weather[0].Description)
		}
		os.Exit(0)
	}

	dg, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch msgfld := strings.Fields(m.Content); strings.ToLower(msgfld[0]){
		case "!weather":
			weather.HandleWeatherRequest(WeatherToken, s,m)
		case  "!def":
			var result string
			if len(msgfld) == 1{
				result = "What do you want to know?"
			} else {
				result = dictionary.GetDefinition(msgfld[1])
			}
			s.ChannelMessageSend(m.ChannelID, result)
		}

	// If the message is "ping" reply with "Pong!"
	/*if strings.ToLower(m.Content) == "!weather" {
		//s.ChannelMessageSend(m.ChannelID, getWeather(4990729))
		weather.HandleWeatherRequest(WeatherToken,s,m)
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}*/
}

func HandleWeatherRequest(s *discordgo.Session, m *discordgo.MessageCreate) {

	type City struct {
		name  string
		locID int
	}

	var cities = []City{City{"Millennium City and Windsor, Ontario", 4990729}}

	for _,c := range cities {

		w, err := owm.NewCurrent("C", "EN", WeatherToken)

		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID,"I cannot get the weather right now.")
		} else {
			w.CurrentByID(c.locID)
			log.Println(w)
			log.Printf("The current temperature in %s is %v degrees\n", w.Name, w.Main.Temp)
			log.Println(len(w.Weather))
			log.Println(w.Weather[0].Description)
			log.Printf("Wind: Direction %v, Speed %v\n", w.Wind.Deg, w.Wind.Speed)
			log.Printf("Rain: %#v\n", w.Rain)
			log.Printf("Snow: %#v\n", w.Snow)
		var wv = convertWindVelocity(w.Wind.Speed)	
		s.ChannelMessageSend(m.ChannelID, 
		fmt.Sprintf("WCOC Action Weather For %v: %v.\nThe current temperature is %v degrees Celcius, "+
			"%v degrees Farenheit.\nHumidity is %v%%, winds are out of the %v at %.1f kph (%.1f mph)",
			c.name,
			w.Weather[0].Description,
			math.Round(w.Main.Temp),
			math.Round((w.Main.Temp * 1.8) + 32),
			w.Main.Humidity,
			getWindDirection(w.Wind.Deg),
			wv.KpH, wv.MpH))
	
	
	}
}
}

func getWindDirection(dir float64) string{
	var strDirection string
	
	if dir < 0 {
		dir = 360-dir
	}
	switch {
	case dir <27.5:
		strDirection =  "north"
	case dir < 72.5:
		strDirection = "north east"
	case dir <  117.5:
		strDirection = "east"
	case dir < 162.5:
		strDirection = "south east"
	case dir <  207.5:
		strDirection = "south"
	case dir <  252.5:
		strDirection = "south west"
	case dir < 297.5:
		strDirection = "west"
	case dir  < 342.5:
		strDirection = "north west"
	case dir <= 360:
		strDirection = "north"
	default:
		strDirection = "Some Direction"
		log.Printf("Direction = %v\n", dir)
	}

	return strDirection
}
/* 
   Wind velocity comes in Meters per Second. This converts to Kilometers
   per hour and Miles per hour
*/
func convertWindVelocity(MpS float64) (converted WindVelocity){
	converted.KpH = MpS*3600/1000
	converted.MpH = converted.KpH*.621371
	return
}
