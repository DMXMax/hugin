package command

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/DMXMax/hugin/util"
	"github.com/bwmarrin/discordgo"
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
	Op: func(s *discordgo.Session, m *discordgo.MessageCreate, i interface{}) (map[string]string, error) {
		params, ok := i.(map[string]string)
		if ok {
			if seed == 0 {
				seed = time.Now().UnixNano()
				rand.Seed(seed)
				log.Printf("Randomizer seeded with %v\n", seed)
			}
			result := processRoll(strings.Fields(m.Content))
			log.Printf("%s rolls dice %#v\n", util.GetAuthorInfo(s, m), result)
			if result.NeedHelp == true {
				showHelp(s, m)
			} else {
				showResult(util.GetNickname(m), result, s, m, params)
			}
			return nil, nil
		} else {
			return nil, nil
		}

	},
}

func showResult(name string, res RollResult, s *discordgo.Session, m *discordgo.MessageCreate,
	params map[string]string) {
	if len(res.Note) == 0 {
		res.Note = "rolls"
	}
	msg := fmt.Sprintf("**%s** *%s*  (%s) %+d   %s: %+d	 **Total: %+d**",
		name, res.Note, getRollString(res.Rolls), params["skillName"], res.RollTotal(), res.ApproachValue, res.GrandTotal())
	s.ChannelMessageSend(m.ChannelID, msg)
}

func processRoll(rollReq []string, m map[string]string) RollResult {
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

	return result
}
func showHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "Are you trying to roll dice? The Dice command looks like this:" +
		"```/r [Approach Value] [Note]```" +
		"If you don't supply an approach value, your value will be Mediocre(0), and the rest of the command string becomes the note."
	s.ChannelMessageSend(m.ChannelID, msg)
}

func getRollString(roll [4]int) (result string) {
	fig := [3]rune{'\u2296', '\u2299', '\u2295'}
	buf := [4]rune{}
	for i, y := range roll {
		buf[i] = fig[y+1]
	}
	result = string(buf[:])
	return
}
