package model

import (
	"fmt"

	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//Event holds all the necissary information for a particular event
type Event struct {
	ID        string   `json:"-"`
	Name      string   `json:"name,omitempty"`
	Date      string   `json:"date,omitempty"`
	Length    string   `json:"length,omitempty"`
	TimeStart []string `json:"time-start,omitempty"`
	TimeEnd   []string `json:"time-end,omitempty"`
}

//Add insert the event into Firebase. If it is unable to then it
//returns an error.
func (e *Event) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("events").Push(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("Error occured while attempting to add the event\n%v", err)
	}
	e.ID = respRef.Key

	return respRef, nil
}
