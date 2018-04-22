package model

import (
	"log"
	"os"

	"firebase.google.com/go/db"

	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var (
	ref *db.Ref
)

func init() {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://shika-app.firebaseio.com/",
	}
	wd, err := os.Getwd() // Get the working directory
	if err != nil {
		log.Println("Could not get current working directory")
	}
	opt := option.WithCredentialsFile(wd + "/../creds/firebase-adminsdk.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Printf("Could not initialize Firebase App: %v", err)
	}
	dbClient, err := app.Database(ctx)
	if err != nil {
		log.Printf("Could not initialize database client: %v", err)
	}

	ref = dbClient.NewRef("")
}
