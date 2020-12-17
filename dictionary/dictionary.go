package dictionary

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"github.com/DMXMax/hugin/data"
)


func HandleDictionaryRequest(words []string, s *dg.Session,
	m *dg.MessageCreate) {
	var (
	//	def Definition
	)
	ctx := context.Background()

	if client, err := datastore.NewClient(ctx, "hugin-00001"); err == nil {
		q := datastore.NewQuery("dictEntry").Filter("word =", strings.
			ToLower(words[0]))
		t := client.Run(ctx, q)
		def := data.DictionaryEntry{}
		if key, err := t.Next(&def); err == nil {
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("%s: %s", def.Display, def.Definition))
			log.Printf("%s: %s, [%s]\n", def.Display, def.Definition, key)
		} else {
			if err == iterator.Done {
				s.ChannelMessageSend(m.ChannelID,
					fmt.Sprintf("Don't know what %v is.", words[0]))
			} else {
				log.Println("Error getting next." + err.Error())
			}
		}
	} else {
		log.Printf("no client" + err.Error())
	}
}
