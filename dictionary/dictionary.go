package dictionary

import (
	"fmt"
	"strings"
	"log"
	pr "github.com/DMXMax/noppa/dictionary/propreader"
	dg "github.com/bwmarrin/discordgo"
)

var dict pr.Dictionary

func getDefinition(word string) string {
	var result string
	var err error
	if len(dict) == 0 {
		fmt.Println("loading dictionary")
		dict, err = pr.ReadDictionaryFile("/home/glen_clarkson_gmail_com/tla.txt")
	}
	if err != nil {
		result = "Can't load dictionary"
	} else {
		if len(dict) == 0 {
			result = "Something is Wrong."
		} else {
			idx := strings.ToLower(word)
			def := dict[idx]
			if def == "" {
				result = fmt.Sprintf("No definiton for %s\n", word)
			} else {
				result = fmt.Sprintf("%s: %s\n", word, dict[idx])
			}
		}
	}
	return result
}

//func messageCreate(s *dg.Session, m *dg.MessageCreate) {
func HandleDictionaryRequest(words []string, s *dg.Session, m *dg.MessageCreate) {

	var result string
	if len(words) == 0 {
		result = "What do you want to know?"
//		testComplexSend(s,m)
	} else {
		result = getDefinition(words[0])
	}
	s.ChannelMessageSend(m.ChannelID, result)

}

func HandleDictionaryReset(s *dg.Session, m *dg.MessageCreate){
	if m.Author.Username == "anonymaustrap" && m.Author.Discriminator == "0551" {
		dict = pr.Dictionary{}
		s.ChannelMessageSend(m.ChannelID, "Dictionary Reload")
	}
}		

func testComplexSend(s *dg.Session, m *dg.MessageCreate){
	field1 := dg.MessageEmbedField{"Name1", "Value1", true}
	field2 := dg.MessageEmbedField{"Name2", "Value2", false}

	fields := []*dg.MessageEmbedField{&field1, &field2}
	video := dg.MessageEmbedVideo{
		URL: "https://www.youtube.com/watch?v=KcmiqQ9NpPE&list=PLC430F6A783A88697",
		Width: 10,
		Height: 10,
	}
	image := dg.MessageEmbedImage{
		URL: "https://upload.wikimedia.org/wikipedia/commons/c/c5/StudioFibonacci_Cartoon_Bat.png",
		Width:10,
		Height:10,
	}

	author := dg.MessageEmbedAuthor{Name: "BobbyJoe"}
	provider := dg.MessageEmbedProvider{Name: "ProviderJoe"}

	embed := dg.MessageEmbed{
		URL:"https://www.youtube.com/watch?v=KcmiqQ9NpPE&list=PLC430F6A783A88697",
		Title: "A Title!",
		Color: 0,
		Description: "A Description",
		Fields: fields,
		Video: &video,
		Author: &author,
		Provider: &provider,
		Image: &image,
	}
	msgSend := dg.MessageSend{
	Content: "Content",
	Embed: &embed,	
	}
	if msg, err :=s.ChannelMessageSendComplex(m.ChannelID, &msgSend); err!=nil{
		log.Printf("Error %v\n",err)
	} else {
		log.Printf("Message %v\n", msg)
	}
}
