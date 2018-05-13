package model

import (
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

type VisitorInterface interface {
	Add(*Visitor) (*db.Ref, error)
	Remove(string) error
}

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

type visitorInterface struct{ VisitorInterface }

//Add inserts a visitor and his/her dog(s) into Firebase.
//The user visiting must alread exist in Firebase along
//with the users dog(s)
func (vi visitorInterface) Add(v *Visitor) (*db.Ref, error) {
	ctx := context.Background()
	ref := client.NewRef("visitors")

	respRef, err := ref.Push(ctx, v)
	if err != nil {
		return nil, err
	}
	v.ID = respRef.Key

	return respRef, nil
}

//Remove deletes the visitor from Firebase
func (vi visitorInterface) Remove(vID string) error {
	ctx := context.Background()
	ref := client.NewRef("visitors")

	err := ref.Child(vID).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
