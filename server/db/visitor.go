package db

import (
	"context"

	"github.com/williabk198/shika_app/server/model"
)

type VisitorInterface interface {
	Add(*model.Visitor) (string, error)
	Get(string) (*model.Visitor, error)
	Remove(string) error
}

type Visitor struct{ VisitorInterface }

//Add inserts a visitor and his/her dog(s) into Firebase.
//The user visiting must alread exist in Firebase along
//with the users dog(s)
func (v Visitor) Add(visitor *model.Visitor) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("visitors")

	respRef, err := ref.Push(ctx, visitor)
	if err != nil {
		return "", err
	}
	visitor.ID = respRef.Key

	return respRef.Key, nil
}

//Remove deletes the visitor from Firebase
func (v Visitor) Remove(vID string) error {
	ctx := context.Background()
	ref := client.NewRef("visitors")

	err := ref.Child(vID).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
