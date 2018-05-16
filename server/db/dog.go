package db

import (
	"context"

	"github.com/williabk198/shika_app/server/model"
)

type DogInterface interface {
	Add(*model.Dog) (string, error)
	Get(string) (*model.Dog, error)
}

type Dog struct{ DogInterface }

//Add inserts a the dog into Firebase and associates
//it to the given owner via the userKey.
func (d Dog) Add(dog *model.Dog) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("dogs")

	respRef, err := ref.Push(ctx, dog)
	if err != nil {
		return "", err
	}
	dog.ID = respRef.Key

	return respRef.Key, nil
}

func (d Dog) Get(dogKey string) (*model.Dog, error) {
	ctx := context.Background()
	ref := client.NewRef("dogs")

	dog := new(model.Dog)
	err := ref.Get(ctx, dog)
	if err != nil {
		return nil, err
	}
	return dog, nil
}
