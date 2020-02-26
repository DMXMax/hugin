package user
import (
	"log"
	"context"
	"cloud.google.com/go/datastore"
)

type User struct{

	Key	*datastore.Key	`datastore:"__key__"`
	Scopes 	[]string	`datastore:"scopes,noindex"`
}

func AddUser(ctx context.Context, client *datastore.Client, userID string,
	user *User) {
	k := datastore.NameKey("user", userID, nil)
	if res, err := client.Put(ctx, k, user); err == nil {
		log.Printf("Added. Result = %v\n", res)
	} else {
		log.Printf("Failed to Add. Error = %v\n", err)
	}
}
func SaveUser( ctx context.Context, client *datastore.Client, u *User) error {
	if res, err := client.Put(ctx,u.Key, u); err == nil {
		log.Printf("Saved: %v\n", res)
		return nil
	} else {
		return err
	}
}
func GetUser(ctx context.Context, client *datastore.Client, userID string) *User {
	k := datastore.NameKey("user", userID, nil)
	u := new(User)
	if err := client.Get(ctx, k, u); err == nil {
	} else {
		log.Printf("Failed to Retrieve: %v\n", err)
		//This creates a generic user. This user can only run commands attributed to the 'any' scope
		u.Key = k
		u.Scopes = []string{"any"}
		if err := SaveUser( ctx, client, u); err == nil{
			log.Printf("Added new user %v\n", u)
		} else {
			log.Printf("Failed to save user %v\n", err)
		}
	}
	log.Printf("User retrieved: %#v\n", u)
	return u
}


func (u *User) HasScope(scope string) bool {
	for _, v := range u.Scopes {
		if v == "admin" || v == scope || "any" == scope {
			return true
		}
	}
	return false
}
