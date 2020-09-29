package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	msg "github.com/DMXMax/hugin/msghandler"
	"github.com/DMXMax/hugin/weather"
	"github.com/DMXMax/hugin/util"
	"github.com/bwmarrin/discordgo"
)

const noppa_pie = "projects/131909995516/secrets/noppa-pie/versions/1"

// Variables used for command line parameters
var (
	DiscordToken string
)

func main() {

	if err := weather.Init(); err != nil {
		log.Fatalln("Error intializing weather service", err)
		os.Exit(1)
	}
	
	DiscordToken, err := util.GetSecret(noppa_pie)

	if err != nil{
		fmt.Println("No token")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(msg.HandleMessageCreate)
	dg.AddHandler(msg.HandleMessageUpdate)
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
	log.Println("Shutting Down")
	dg.Close()
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
