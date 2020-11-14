package util

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

func GetAuthorInfo(s *discordgo.Session, m *discordgo.MessageCreate) string {
	var prefix string
	if m.GuildID != "" {
		if g, err := s.Guild(m.GuildID); err == nil {
			prefix = fmt.Sprintf("%s:%s", g.Name, m.Author.Username)
		} else {
			log.Println("Error, HandleMessageCreate, retrieving guild", err)
		}
	} else {
		prefix = fmt.Sprintf("%s", m.Author.Username)
	}
	return prefix
}

func GetNickname(m *discordgo.MessageCreate) string {
	name := "Nobody"
	if m.Member != nil {
		name = m.Member.Nick
		if name == ""{
			if m.Author != nil{
				name = m.Author.Username
			}
		}
	} else {
		if m.Author != nil {
			name = m.Author.Username
		}
	}
	return name
}

func GetSecret(version string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)

	if err == nil {
		req := &secretmanagerpb.AccessSecretVersionRequest{
			Name: version,
		}
		result, err := client.AccessSecretVersion(ctx, req)
		if err == nil {
			return string(result.Payload.Data), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}
