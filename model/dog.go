package model

import (
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//Dog holds all the information in regards to a users pet
type Dog struct {
	Name      string `json:"name,omitempty"`
	Breed     string `json:"breedd,omitempty"`
	ImageURL  string `json:"image-url,omitempty"`
	CheckedIn bool   `json:"checked-in"`
}

//Add inserts a the dog into Firebase and associates
//it to the given owner via the userKey.
func (d *Dog) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("dogs").Push(ctx, d)
	if err != nil {
		return nil, err
	}

	return respRef, nil
}

//UpdateCheckIn updates whether a dog has been checked into the park or not
func (d *Dog) UpdateCheckIn(dogRef *db.Ref, isCheckedIn bool) error {
	ctx := context.Background()

	d.CheckedIn = isCheckedIn
	dogKey := dogRef.Key

	dMap := map[string]interface{}{dogKey: d}
	err := dogRef.Parent().Update(ctx, dMap)
	if err != nil {
		return err
	}

	return nil
}
