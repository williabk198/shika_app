package model

import (
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//Dog holds all the information in regards to a users pet
type Dog struct {
	ID       string `json:"-"`
	Name     string `json:"name,omitempty"`
	Breed    string `json:"breed,omitempty"`
	Sex      string `json:"sex,omitempty"`
	ImageURL string `json:"image-url,omitempty"`
}

//Add inserts a the dog into Firebase and associates
//it to the given owner via the userKey.
func (d *Dog) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("dogs").Push(ctx, d)
	if err != nil {
		return nil, err
	}
	d.ID = respRef.Key

	return respRef, nil
}
