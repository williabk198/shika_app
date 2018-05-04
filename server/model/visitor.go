package model

import (
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//Visitor holds which user and which dog(s)
//the user has checked in along with the area
//that the user checked in to.
type Visitor struct {
	ID          string   `json:"-"`
	UserKey     string   `json:"user"`
	DogKeys     []string `json:"dogs"`
	CheckInTime string   `json:"check-in"`
	ParkArea    string   `json:"park-area"`
}

//Add inserts a visitor and his/her dog(s) into Firebase.
//The user visiting must alread exist in Firebase along
//with the users dog(s)
func (v *Visitor) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("visitors").Push(ctx, v)
	if err != nil {
		return nil, err
	}
	v.ID = respRef.Key

	return respRef, nil
}

//Remove deletes the visitor from Firebase
func (v *Visitor) Remove() error {
	ctx := context.Background()

	err := ref.Child("visitors/" + v.ID).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
