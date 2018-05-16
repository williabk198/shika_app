package db

import (
	"context"
	"fmt"

	"github.com/williabk198/shika_app/server/model"
)

type UserInterface interface {
	Add(*model.User) (string, error)
	AddDog(string, string) error
	Get(string) (*model.User, error)
}

type User struct{ UserInterface }

//Add inserts the user into Firebase
func (u User) Add(user *model.User) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("users")

	respRef, err := ref.Push(ctx, user)
	if err != nil {
		return "", err
	}
	user.ID = respRef.Key

	return respRef.Key, nil
}

//AddDogKey adds a key associated to Dog already existing in Firebase
//and updates the user.
func (u User) AddDogKey(userKey, dogKey string) error {
	ctx := context.Background()
	ref := client.NewRef("")

	user, err := u.Get(userKey)
	if err != nil {
		return fmt.Errorf(
			"Error could not retrieve user with ID %s\n%v",
			userKey, err,
		)
	}

	result, err := ref.Child("dogs").OrderByKey().
		EqualTo(dogKey).GetOrdered(ctx)

	if err != nil {
		return fmt.Errorf("Error attempting to check if dog "+
			"exists in the database\n%v", err)
	}
	if len(result) == 0 {
		//TODO: Make this error into its own type
		return fmt.Errorf("Error associating to dog to its user. " +
			"The dog does not exist in the database")
	}

	user.Dogs = append(user.Dogs, result[0].Key())

	uMap := map[string]interface{}{userKey: user}
	err = ref.Child("users").Update(ctx, uMap)
	if err != nil {
		return err
	}
	return nil
}

func (u User) Get(userKey string) (*model.User, error) {
	ctx := context.Background()
	ref := client.NewRef("users")

	user := new(model.User)
	err := ref.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
