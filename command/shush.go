package command

import (

	"github.com/bwmarrin/discordgo"
)

var ShushCommand Command = Command{
	Name : "Shush",
	Scope : "admin",
	Op : func(s *discordgo.Session, m *discordgo.MessageCreate) (map[string]string,error) { 
		return nil,nil
	},
}
