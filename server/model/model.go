package model

import (
	"log"

	"firebase.google.com/go/db"

	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var (
	Users    UserInterface    = userInterface{}
	Dogs     DogInterface     = dogInterface{}
	Events   EventInterface   = eventInterface{}
	Visitors VisitorInterface = visitorInterface{}

	client *db.Client
)

func init() {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://shika-app.firebaseio.com/",
	}
	opt := option.WithCredentialsFile("../creds/firebase-adminsdk.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Printf("Could not initialize Firebase App: %v", err)
	}
	client, err = app.Database(ctx)
	if err != nil {
		log.Printf("Could not initialize database client: %v", err)
	}
}
