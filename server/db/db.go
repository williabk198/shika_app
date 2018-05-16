package db

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var (
	Users    UserInterface
	Dogs     DogInterface
	Events   EventInterface
	Visitors VisitorInterface

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
