package msghandler 

import (
	"log"
	"fmt"
	"strings"
	"github.com/DMXMax/noppa/dictionary"
	"github.com/DMXMax/noppa/weather"
	"github.com/bwmarrin/discordgo"
	"github.com/DMXMax/noppa/whois"
	"github.com/DMXMax/noppa/command"
)

var shushFlag bool = false

// message is created on any channel that the autenticated bot has access to.
func HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if shushFlag && m.Author.ID != "427628065720500239"{
		return
	}
	log.Printf("Guild %v\n", m.GuildID)
	if g,err := s.Guild(m.GuildID); err == nil {
		log.Printf("Guild = %v\n", g.Name)
	} else {
		log.Println(err)
	}

	msgflds := strings.Fields(m.Content)
	if len(msgflds) > 0 {
		switch strings.ToLower(msgflds[0]) {
		case "!weather":
			//weather.HandleWeatherRequest(s, m)
			_ = weather.WeatherCommand.Call(s,m)
		case "!def":
			dictionary.HandleDictionaryRequest(msgflds[1:], s, m)
		case "!ping":
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("Pong, %s", m.Author.Mention()))
		case "!whois":
			whois.HandleRequest(msgflds[1:], s,m)
		case "!shush":
			if err := command.ShushCommand.Call(s,m); err == nil{
				shushFlag = !shushFlag
				if shushFlag{
					s.ChannelMessageSend(m.ChannelID, "Entering Shush Mode")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Listening")
				}
			}
		}
	}
}



