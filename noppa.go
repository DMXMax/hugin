package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DMXMax/noppa/dictionary"
	"github.com/DMXMax/noppa/weather"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	DiscordToken string
	WeatherToken string
	weatherOnly  *bool
)

func init() {

	weatherOnly = flag.Bool("weather", false, "Test Weather Only")
	flag.Parse()
}

func main() {

	if WeatherToken = os.Getenv("NOPPA_WEATHER"); WeatherToken == "" {
		fmt.Println("No Weather Token")
		os.Exit(1)
	}

	if DiscordToken = os.Getenv("NOPPA_PIE"); DiscordToken == "" {
		fmt.Println("No token")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(ready)

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
	fmt.Printf("Message from %v#%v\n", m.Author.Username, m.Author.Discriminator)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch msgflds := strings.Fields(m.Content); strings.ToLower(msgflds[0]) {
	case "!weather":
		weather.HandleWeatherRequest(WeatherToken, s, m)
	case "!def":
		dictionary.HandleDictionaryRequest(msgflds[1:], s, m)
	case "!reload":
		dictionary.HandleDictionaryReset(s,m)
	}
}
func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	sd := discordgo.UpdateStatusData{
		nil,
		nil,
		false,
		"Attentive",
	}

	if err := s.UpdateStatusComplex(sd); err != nil {
		log.Printf("Error setting status %v\n", err)
	} else {
		log.Println("Status Set")
	}

}
