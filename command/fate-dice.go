package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RollResult struct {
	ApproachValue int
	Note          string
	Rolls         [4]int
	NeedHelp      bool
}

func (r RollResult) RollTotal() int {
	return r.Rolls[0] + r.Rolls[1] + r.Rolls[2] + r.Rolls[3]
}
func (r RollResult) GrandTotal() int {
	return r.RollTotal() + r.ApproachValue
}

var seed int64

var FateDiceCommand Command = Command{
	Name:  "fate-dice",
	Scope: "any",
	Op: func(s *discordgo.Session, m *discordgo.MessageCreate) (map[string]string, error) {
		if seed == 0 {
			seed = time.Now().UnixNano()
			rand.Seed(seed)
			log.Printf("Randomizer seeded with %v\n", seed)
		}
		result := processRoll(strings.Fields(m.Content))
		if result.NeedHelp == true {
			showHelp(s, m)
		} else {
			showResult(getUsableName(m), result, s, m)
		}
		return nil, nil

	},
}

func showResult(name string, res RollResult, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(res.Note) == 0{
		res.Note = "rolls"
	}
	msg := fmt.Sprintf("**%s** *%s*  (%+d,%+d,%+d,%+d) %+d   Approach: %+d	 **Total: %+d**",
		name, res.Note,  res.Rolls[0], res.Rolls[1], 
		res.Rolls[2], res.Rolls[3],res.RollTotal(),  res.ApproachValue, res.GrandTotal())
	s.ChannelMessageSend(m.ChannelID, msg)
}

func processRoll(rollReq []string) RollResult {
	note := ""
	approach := 0
	needhelp := false

	if len(rollReq) > 1 {
		firstArg := strings.ToLower(rollReq[1])
		if firstArg == "-h" || firstArg == "-?" || firstArg == "/?" || firstArg == "/h" {
			needhelp = true
		} else {
			if val, err := strconv.Atoi(firstArg); err == nil {
				approach = val
				note = strings.Join(rollReq[2:], " ")
			} else {
				note = strings.Join(rollReq[1:], " ")
			}
		}
	}

	result := RollResult{
		ApproachValue: approach,
		Note:          note,
		Rolls:         [4]int{rand.Intn(3) - 1, rand.Intn(3) - 1, rand.Intn(3) - 1, rand.Intn(3) - 1},
		NeedHelp:      needhelp,
	}

	log.Printf("%v\n", rollReq)

	return result
}
func showHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg :="Are you trying to roll dice? The Dice command looks like this:"+
	"```/r [Approach Value] [Note]```"+
	"If you don't supply an approach value, your value will be Mediocre(0), and the rest of the command string becomes the note."
	s.ChannelMessageSend(m.ChannelID, msg)
}


func getUsableName(m *discordgo.MessageCreate) string {
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
