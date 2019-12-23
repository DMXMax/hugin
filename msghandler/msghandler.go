package msghandler 

import (
	"fmt"
	"strings"
	"github.com/DMXMax/noppa/dictionary"
	"github.com/DMXMax/noppa/weather"
	"github.com/bwmarrin/discordgo"
)

// message is created on any channel that the autenticated bot has access to.
func HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	msgflds := strings.Fields(m.Content)
	if len(msgflds) > 0 {
		switch strings.ToLower(msgflds[0]) {
		case "!weather":
			weather.HandleWeatherRequest(s, m)
		case "!def":
			dictionary.HandleDictionaryRequest(msgflds[1:], s, m)
		case "!ping":
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("Pong, %s", m.Author.Mention()))
		}
	}
}
