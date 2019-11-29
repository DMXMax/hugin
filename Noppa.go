package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ramsgoli/openweathermap"
)

// Variables used for command line parameters
var (
	Token string
	weatherOnly *bool
)

func init() {

//	flag.StringVar(&Token, "t", "", "Bot Token")
//	flag.Parse()
	weatherOnly = flag.Bool("weather",false,"Test Weather Only")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	//dg, err := discordgo.New("Bot " + Token)
	if Token = os.Getenv("NOPPA_PIE"); Token == ""{
		fmt.Println("No token")
		os.Exit(1)
	}

	if *weatherOnly==true {
		fmt.Println("Testing Weather Only")
		os.Exit(0)
	}

	dg, err := discordgo.New("Bot " + Token)
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
		s.ChannelMessageSend(m.ChannelID, `WCOC Weather: Mostly cold with a chance of light flurries.
	The current temperature is 4 degrees Celcius, 39 degrees Farenheit.`)
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
