package model

import (
	"fmt"

	"golang.org/x/net/context"
)

type EventInterface interface {
	Add(*Event) (string, error)
}

//Event holds all the necissary information for a particular event
type Event struct {
	ID        string   `json:"-"`
	Name      string   `json:"name,omitempty"`
	Date      string   `json:"date,omitempty"`
	Length    string   `json:"length,omitempty"`
	TimeStart []string `json:"time-start,omitempty"`
	TimeEnd   []string `json:"time-end,omitempty"`
}

type eventInterface struct{ EventInterface }

//Add insert the event into Firebase. If it is unable to then it
//returns an error.
func (ei eventInterface) Add(e *Event) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("events")

	respRef, err := ref.Push(ctx, e)
	if err != nil {
		return "", fmt.Errorf("Error occured while attempting to add the event\n%v", err)
	}
	e.ID = respRef.Key

	return respRef.Key, nil
}
