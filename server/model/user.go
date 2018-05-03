package model

import (
	"fmt"
	"time"

	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

//User holds the information that of the person visiting the dog park
type User struct {
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	ImgURL      string    `json:"image,omitempty"`
	Dogs        []string  `json:"dogs,omitempty"`
	CheckedIn   bool      `json:"checked-in"`
	CheckInTime time.Time `json:"checked-in-time,omitempty"`
	Area        string    `json:"area,omitempty"`
}

//Add inserts the user into Firebase
func (u *User) Add() (*db.Ref, error) {
	ctx := context.Background()

	respRef, err := ref.Child("users").Push(ctx, u)
	if err != nil {
		return nil, err
	}

	return respRef, nil
}

//AddDogKey adds a key associated to Dog already existing in Firebase
//and updates the user.
func (u *User) AddDogKey(userKey string, dogKey string) error {
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

	uMap := map[string]interface{}{userKey: u}
	err = ref.Child("users").Update(ctx, uMap)
	if err != nil {
		return err
	}
	return nil
}

//UpdateCheckIn updates the checked in status of the user
func (u *User) UpdateCheckIn(userRef *db.Ref, isCheckedIn bool) error {
	ctx := context.Background()
	u.CheckedIn = isCheckedIn
	u.CheckInTime = time.Now()
	userKey := userRef.Key

	uMap := map[string]interface{}{userKey: u}
	err := ref.Parent().Update(ctx, uMap)
	if err != nil {
		return err
	}
	return nil
}

//SetArea set the area of the dog park in which the user will be in
func (u *User) SetArea(userRef *db.Ref, area string) error {
	ctx := context.Background()

	u.Area = area
	userKey := userRef.Key

	uMap := map[string]interface{}{userKey: u}
	err := ref.Parent().Update(ctx, uMap)
	if err != nil {
		return err
	}
	return nil
}
