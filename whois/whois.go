package whois

import ( 
	"fmt"
	"strings"
	"log"
	dg "github.com/bwmarrin/discordgo"
)

func HandleRequest( entry []string, s *dg.Session, m *dg.MessageCreate) {
	if strings.ToLower(entry[0]) == "hugin"{
		var me dg.MessageEmbed
		var img dg.MessageEmbedImage = dg.MessageEmbedImage{}
		img.URL = "https://storage.googleapis.com/c1f16fc7-37dc-4643-814a-5394690366a9/UNTIL_LOGO_Tiny.png"
		me.Title = "HUGIN"
		me.Footer = &dg.MessageEmbedFooter{Text:"HUGIN Data"}
		me.Description = "HUGIN is an artificial intelligence maintained by UNTIL."
		me.Image=&img
		if _, err := s.ChannelMessageSendEmbed(m.ChannelID,&me); err != nil{
		log.Printf("Error Sending")
		} 
	} else {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I don't know %s.", entry[0]))
	}
}
