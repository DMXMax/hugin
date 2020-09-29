package util

import(
	"fmt"
	"log"
	"github.com/bwmarrin/discordgo"

)

func GetAuthorInfo(s *discordgo.Session, m *discordgo.MessageCreate) string {
	var prefix string
	if m.GuildID != ""{
		if g,err := s.Guild(m.GuildID); err == nil {
			prefix=fmt.Sprintf("%s:%s", g.Name, m.Author.Username)
		} else {
			log.Println("Error, HandleMessageCreate, retrieving guild", err)
		}
	}else{
		prefix = fmt.Sprintf("%s", m.Author.Username)
	}
	return prefix
}

func GetNickname(m *discordgo.MessageCreate) string {
	name := "Nobody"
	if m.Member != nil {
		name = m.Member.Nick
	} else {
		if m.Author != nil {
			name = m.Author.Username
		}
	}
	return name
}
