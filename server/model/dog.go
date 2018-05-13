package model

import (
	"golang.org/x/net/context"
)

type DogInterface interface {
	Add(*Dog) (string, error)
	Get(string) (*Dog, error)
}

//Dog holds all the information in regards to a users pet
type Dog struct {
	ID       string `json:"-"`
	Name     string `json:"name,omitempty"`
	Breed    string `json:"breed,omitempty"`
	Sex      string `json:"sex,omitempty"`
	ImageURL string `json:"image-url,omitempty"`
}

type dogInterface struct{ DogInterface }

//Add inserts a the dog into Firebase and associates
//it to the given owner via the userKey.
func (di dogInterface) Add(d *Dog) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("dogs")

	respRef, err := ref.Push(ctx, d)
	if err != nil {
		return "", err
	}
	d.ID = respRef.Key

	return respRef.Key, nil
}

func (di dogInterface) Get(dogKey string) (*Dog, error) {
	ctx := context.Background()
	ref := client.NewRef("dogs")

	dog := new(Dog)
	err := ref.Get(ctx, dog)
	if err != nil {
		return nil, err
	}
	return dog, nil
}
