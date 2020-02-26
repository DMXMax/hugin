package command

import (

	"github.com/bwmarrin/discordgo"
)

var ShushCommand Command = Command{
	Name : "GetWeather",
	Scope : "admin",
	Op : func(s *discordgo.Session, m *discordgo.MessageCreate) error { 
		return nil
	},
}
