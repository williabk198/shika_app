package db

import (
	"context"

	"github.com/williabk198/shika_app/server/model"
)

type EventInterface interface {
	Add(*model.Event) (string, error)
	Get(string) (*model.Event, error)
}

type Event struct{ EventInterface }

//Add inserts the event into database. If it is unable to then it
//returns an error.
func (e Event) Add(event *model.Event) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("events")

	respRef, err := ref.Push(ctx, event)
	if err != nil {
		return "", err
	}
	event.ID = respRef.Key

	return respRef.Key, nil
}

//Get retrieves an event with the given eventKey from the database
func (e Event) Get(eventKey string) (*model.Event, error) {
	ctx := context.Background()
	ref := client.NewRef("events")

	event := new(model.Event)
	err := ref.Child(eventKey).Get(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
