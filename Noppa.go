package main

import (
	"flag"
	"fmt"
	 "math"
	"os"
	"os/signal"
	"strings"
	"syscall"

	owm "github.com/briandowns/openweathermap"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	DiscordToken string
	WeatherToken string
	weatherOnly  *bool
)

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
		fmt.Println("Testing Weather Only")
		w, err := owm.NewCurrent("C", "EN", WeatherToken)
		if err != nil {
			fmt.Println(err)
		} else {
			w.CurrentByID(4990729)
			fmt.Println(w)
			fmt.Printf("The current temperature in %s is %v degrees\n", w.Name, w.Main.Temp)
			fmt.Println(w.Weather[0].Description)
		}
		os.Exit(0)
	}

	dg, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
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
	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "!weather" {
		//s.ChannelMessageSend(m.ChannelID, getWeather(4990729))
		HandleWeatherRequest(s,m)
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func HandleWeatherRequest(s *discordgo.Session, m *discordgo.MessageCreate) {

	type City struct {
		name  string
		locID int
	}

	var cities = []City{City{"Millennium City", 4990729}, City{"Windsor, Ontario", 6182959}}

	for _,c := range cities {

		w, err := owm.NewCurrent("C", "EN", WeatherToken)

		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID,"I cannot get the weather right now.")
		} else {
			s.ChannelMessageSend(m.ChannelID,"WCOC Action Weather:\n")
			w.CurrentByID(c.locID)
			fmt.Println(w)
			fmt.Printf("The current temperature in %s is %v degrees\n", w.Name, w.Main.Temp)
			fmt.Println(len(w.Weather))
			fmt.Println(w.Weather[0].Description)
		
		s.ChannelMessageSend(m.ChannelID, 
		fmt.Sprintf("For %v: %v.\nThe current temperature is %v degrees Celcius, "+
			"%v degrees Farenheit.",
			c.name,
			w.Weather[0].Description,
			math.Round(w.Main.Temp),
			math.Round((w.Main.Temp * 1.8) + 32)))
	
	
	}
}
}
