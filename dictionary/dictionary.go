package dictionary

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"github.com/DMXMax/noppa/hugindb"
	"log"
)


func HandleDictionaryRequest(words []string, s *dg.Session, m *dg.MessageCreate) {
	var (
		acronym    string
		definition string
	)
	err := hugindb.DB.QueryRow(
		"SELECT acronym, definition FROM TLA where idx = ?", words[0]).
		Scan(&acronym, &definition)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID,
			fmt.Sprintf("I don't know what %v is.", words[0]))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s:  %s", acronym, definition))
	}
}


func testComplexSend(s *dg.Session, m *dg.MessageCreate) {
	field1 := dg.MessageEmbedField{"Name1", "Value1", true}
	field2 := dg.MessageEmbedField{"Name2", "Value2", false}

	fields := []*dg.MessageEmbedField{&field1, &field2}
	video := dg.MessageEmbedVideo{
		URL:    "https://www.youtube.com/watch?v=KcmiqQ9NpPE&list=PLC430F6A783A88697",
		Width:  10,
		Height: 10,
	}
	image := dg.MessageEmbedImage{
		URL:    "https://upload.wikimedia.org/wikipedia/commons/c/c5/StudioFibonacci_Cartoon_Bat.png",
		Width:  10,
		Height: 10,
	}

	author := dg.MessageEmbedAuthor{Name: "BobbyJoe"}
	provider := dg.MessageEmbedProvider{Name: "ProviderJoe"}

	embed := dg.MessageEmbed{
		URL:         "https://www.youtube.com/watch?v=KcmiqQ9NpPE&list=PLC430F6A783A88697",
		Title:       "A Title!",
		Color:       0,
		Description: "A Description",
		Fields:      fields,
		Video:       &video,
		Author:      &author,
		Provider:    &provider,
		Image:       &image,
	}
	msgSend := dg.MessageSend{
		Content: "Content",
		Embed:   &embed,
	}
	if msg, err := s.ChannelMessageSendComplex(m.ChannelID, &msgSend); err != nil {
		log.Printf("Error %v\n", err)
	} else {
		log.Printf("Message %v\n", msg)
	}
}
