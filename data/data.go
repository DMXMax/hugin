package data

import( 
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

type DictionaryEntry struct {
	Word       string	`datastore:"word"`
	Display    string	`datastore:"display,noindex`
	Definition string	`datastore:"definition,noindex`
	Creator	   string	`datastore:"creator"`
}



var DataClient *datastore.Client

func init(){
	if DataClient == nil{
		if d,err := datastore.NewClient(context.Background(), "hugin-00001");err!= nil{
			log.Fatalf("Error %v\n", err)
		} else {

			DataClient = d
		}
	}
}
