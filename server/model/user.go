package model

import (
	"fmt"

	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//User holds the information that of the person visiting the dog park
type User struct {
	ID     string   `json:"-"`
	Name   string   `json:"name,omitempty"`
	Email  string   `json:"email,omitempty"`
	ImgURL string   `json:"image,omitempty"`
	Dogs   []string `json:"dogs,omitempty"`
}

//Add inserts the user into Firebase
func (u *User) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("users").Push(ctx, u)
	if err != nil {
		return nil, err
	}
	u.ID = respRef.Key

	return respRef, nil
}

//AddDogKey adds a key associated to Dog already existing in Firebase
//and updates the user.
func (u *User) AddDogKey(dogKey string) error {
	ctx := context.Background()

	result, err := ref.Child("dogs").OrderByKey().
		EqualTo(dogKey).GetOrdered(ctx)

	if err != nil {
		return fmt.Errorf("Error attempting to check if dog "+
			"exists in database\n%v", err)
	}
	if len(result) == 0 {
		//TODO: Make this error into its own type
		return fmt.Errorf("Error associating to dog to its user. " +
			"The dog does not exist in the database")
	}

	u.Dogs = append(u.Dogs, result[0].Key())

	uMap := map[string]interface{}{u.ID: u}
	err = ref.Child("users").Update(ctx, uMap)
	if err != nil {
		return err
	}
	return nil
}
