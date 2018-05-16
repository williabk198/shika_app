package db

import (
	"context"
	"fmt"

	"github.com/williabk198/shika_app/server/model"
)

type EventInterface interface {
	Add(*model.Event) (string, error)
}

type Event struct{ EventInterface }

//Add insert the event into Firebase. If it is unable to then it
//returns an error.
func (e Event) Add(event *model.Event) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("events")

	respRef, err := ref.Push(ctx, event)
	if err != nil {
		return "", fmt.Errorf("Error occured while attempting to add the event\n%v", err)
	}
	event.ID = respRef.Key

	return respRef.Key, nil
}
