package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DMXMax/noppa/hugindb"
	"github.com/DMXMax/noppa/weather"
	"github.com/bwmarrin/discordgo"
	msg "github.com/DMXMax/noppa/msghandler"
)

// Variables used for command line parameters
var (
	DiscordToken string
)

func main() {

	if err := weather.Init(); err != nil{
		log.Fatalln("Error intializing weather service", err)
		os.Exit(1)
	}
	if err := hugindb.Init(); err != nil{
		log.Fatalln("Error initializing database", err)
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
	dg.AddHandler(msg.HandleMessageCreate)
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
	hugindb.DB.Close()
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
