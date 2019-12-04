package dictionary

import (
	pr "github.com/DMXMax/noppa/dictionary/propreader" 
	"fmt"
	"strings"
)

var dict pr.Dictionary

func GetDefinition(word string) string {
	var result string
	var err error
	if len(dict) == 0 {
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
