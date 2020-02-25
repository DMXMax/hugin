package data

import( 
_	"cloud.google.com/go/datastore"
)

type DictionaryEntry struct {
	Word       string	`datastore:"word"`
	Display    string	`datastore:"display,noindex`
	Definition string	`datastore:"definition,noindex`
}


type User struct{
	UserID string
	Scopes []string
}
