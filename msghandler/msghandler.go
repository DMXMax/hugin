package msghandler

import (
	"fmt"
	"log"
	"strings"

	"github.com/DMXMax/hugin/command"
	"github.com/DMXMax/hugin/dictionary"
	"github.com/DMXMax/hugin/util"
	"github.com/DMXMax/hugin/weather"
	"github.com/DMXMax/hugin/whois"
	"github.com/bwmarrin/discordgo"
)

var shushFlag bool = false

// message is created on any channel that the autenticated bot has access to.
func HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//ignore message to self and messages where the author cannot be determined
	if m.Author == nil || (m.Author.ID == s.State.User.ID || m.Author.ID == "") {
		return
	}

	if shushFlag && m.Author.ID != "427628065720500239" {
		return
	}

	msgflds := strings.Fields(m.Content)
	if len(msgflds) > 0 {
		switch strings.ToLower(msgflds[0]) {
		case "/weather":
			//weather.HandleWeatherRequest(s, m)
			_, _ = weather.WeatherCommand.Call(s, m, nil)
			log.Printf("%s checked the weather", util.GetAuthorInfo(s, m))
		case "/def":
			log.Printf("%s got a definition", util.GetAuthorInfo(s, m))
			if len(msgflds) > 1 {
				dictionary.HandleDictionaryRequest(msgflds[1:], s, m)
			} else {
				s.ChannelMessageSend(m.ChannelID, ":confused:")
			}
		case "/ping":
			log.Printf("%s sent a Ping", util.GetAuthorInfo(s, m))
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("Pong, %s", m.Author.Mention()))
		case "/whois":
			if len(msgflds) > 1 {
				whois.HandleRequest(msgflds[1:], s, m)
			} else {
				s.ChannelMessageSend(m.ChannelID, ":shrug:")
			}
		case "/shush": //shush returns nothing and we're only looking for an error
			if _, err := command.ShushCommand.Call(s, m, nil); err == nil {
				shushFlag = !shushFlag
				if shushFlag {
					s.ChannelMessageSend(m.ChannelID, "Entering Shush Mode")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Listening")
				}
			}
		case "/rfae":
			if _, err := command.FateDiceCommand.Call(s, m, map[string]string{
				"skillName": "Approach",
			}); err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, ":sad:")
			}
		}
	}
}

func HandleMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	HandleMessageCreate(s, &discordgo.MessageCreate{Message: m.Message})
}
