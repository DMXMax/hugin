package dictionary

import (
	 "fmt"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
	"context"
	"strings"
)

type NewDef struct {
	Word string
	Display string
	Definition string 
}

/*type Definition struct {
	acronym    string
	definition string
	adv_flag   bool
	adv_url    sql.NullString
	adv_image  sql.NullString
	adv_text   sql.NullString
}*/

const qry = "SELECT acronym, definition, adv_flag, adv_url, adv_image, " +
	"adv_text FROM TLA WHERE idx = ?"

func HandleDictionaryRequest(words []string, s *dg.Session,
	m *dg.MessageCreate) {
	var (
	//	def Definition
	)
	ctx := context.Background()


	if client, err := datastore.NewClient(ctx, "hugin-00001");err == nil{
		log.Println("Got a client")
		q := datastore.NewQuery("dictEntry").Filter("word =", strings.
			ToLower(words[0]))
		t := client.Run(ctx, q)
		def :=  NewDef{}
		if key,err := t.Next(&def); err == nil{
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
  	}else{
		  log.Printf("no client" + err.Error())
	  }


	/*} else {
		log.Printf("Error getting client %s\n", err.Error()
	}


	if len(words) > 0 {
		err := hugindb.DB.QueryRow(qry, words[0]).Scan(&def.acronym,
			&def.definition, &def.adv_flag, &def.adv_url,
			&def.adv_image, &def.adv_text)
		if err != nil {
			log.Printf("%s\n", err)
			s.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("I don't know what %v is.", words[0]))
		} else {
			log.Printf("Definition: %s: Has Advanced Info: %t\n", words[0], def.adv_flag)
			if def.adv_flag {
				ComplexDefSend(s, m, &def)
			} else {
				s.ChannelMessageSend(m.ChannelID,
					fmt.Sprintf("%s:  %s", def.acronym, def.definition))
			}
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "What do you want to know? Type !def <word>")
	}*/
}

/*func ComplexDefSend(s *dg.Session, m *dg.MessageCreate, def *Definition) {

	embed := dg.MessageEmbed{
		Title:       def.acronym,
		Description: def.definition,
	}
	if def.adv_image.Valid {
		embed.Image = &dg.MessageEmbedImage{
			URL: def.adv_image.String,
		}
	}

	if def.adv_url.Valid {
		embed.URL = def.adv_url.String
	}

	if def.adv_text.Valid {
		embed.Description = def.adv_text.String
	}

	msgSend := dg.MessageSend{
		Embed:   &embed,
	}
	if _, err := s.ChannelMessageSendComplex(m.ChannelID, &msgSend); err != nil {
		log.Printf("Error %v\n", err)
	}
}*/
