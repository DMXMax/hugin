package msghandler 

import (
	"log"
	"fmt"
	"strings"
	"github.com/DMXMax/hugin/dictionary"
	"github.com/DMXMax/hugin/weather"
	"github.com/bwmarrin/discordgo"
	"github.com/DMXMax/hugin/whois"
	"github.com/DMXMax/hugin/command"
)

var shushFlag bool = false

// message is created on any channel that the autenticated bot has access to.
func HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//ignore message to self and messages where the author cannot be determined
	if m.Author == nil{
		log.Println("Message sends nil author")
		log.Println(m.Content)
	}
	if m.Author == nil || (m.Author.ID == s.State.User.ID || m.Author.ID == "") {
		return
	}

	if shushFlag && m.Author.ID != "427628065720500239"{
		return
	}
	if( m.GuildID != ""){
		if g,err := s.Guild(m.GuildID); err == nil {
			log.Printf("Guild = %v\n", g.Name)
		} else {
			log.Println("Error, HandleMessageCreate, retrieving guild", err)

		}
	}else{
		log.Println("No Guild ID availabile. Private Message?")
	}

	msgflds := strings.Fields(m.Content)
	if len(msgflds) > 0 {
		switch strings.ToLower(msgflds[0]) {
		case "/weather":
			//weather.HandleWeatherRequest(s, m)
			_,_ = weather.WeatherCommand.Call(s,m)
		case "/def":
			if  len(msgflds) > 1{
				dictionary.HandleDictionaryRequest(msgflds[1:], s, m)
			} else {
				s.ChannelMessageSend(m.ChannelID, ":confused:")
			}
		case "/ping":
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("Pong, %s", m.Author.Mention()))
		case "/whois":
			if len(msgflds) > 1 {
				whois.HandleRequest(msgflds[1:], s,m)
			}else{
				s.ChannelMessageSend(m.ChannelID,":shrug:")
			}
		case "/shush"://shush returns nothing and we're only looking for an error
			if _,err := command.ShushCommand.Call(s,m); err == nil{
				shushFlag = !shushFlag
				if shushFlag{
					s.ChannelMessageSend(m.ChannelID, "Entering Shush Mode")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Listening")
				}
			}
		case "/r":
			if _, err := command.FateDiceCommand.Call(s,m); err != nil{
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, ":sad:")
			}
		}
	}
}


func HandleMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate){
	HandleMessageCreate(s, &discordgo.MessageCreate{Message:m.Message})
}
