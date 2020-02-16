package dictionary

import (
	"database/sql"
	"fmt"
	"github.com/DMXMax/noppa/hugindb"
	dg "github.com/bwmarrin/discordgo"
	"log"
)

type Definition struct {
	acronym    string
	definition string
	adv_flag   bool
	adv_url    sql.NullString
	adv_image  sql.NullString
	adv_text   sql.NullString
}

const qry = "SELECT acronym, definition, adv_flag, adv_url, adv_image, " +
	"adv_text FROM TLA WHERE idx = ?"

func HandleDictionaryRequest(words []string, s *dg.Session,
	m *dg.MessageCreate) {
	var (
		def Definition
	)
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
	}
}

func ComplexDefSend(s *dg.Session, m *dg.MessageCreate, def *Definition) {

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
}
