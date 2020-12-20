package command

import (
	"context"

	"github.com/DMXMax/hugin/data"
	"github.com/DMXMax/hugin/user"
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name  string
	Scope string
	Op    func(*discordgo.Session, *discordgo.MessageCreate, interface{}) (map[string]string, error)
}

//Haha, we have everything we need to look up the user here. So lets do it here

func (c Command) Call(s *discordgo.Session, msg *discordgo.MessageCreate, i interface{}) (map[string]string, error) {
	//Call getUser here. If no scopes for the user, create a generic user with no scopes
	u := user.GetUser(context.Background(), data.DataClient, msg.Author.ID)
	if u.HasScope(c.Scope) == true {
		return c.Op(s, msg, i)

	} else {
		return nil, NewScopeDeniedError(c.Name, c.Scope)
	}

}

type ScopeDeniedError struct {
}

func (s *ScopeDeniedError) Error() string { return "Scope Denied" }

func NewScopeDeniedError(c string, sr string) error {
	return new(ScopeDeniedError)
}
