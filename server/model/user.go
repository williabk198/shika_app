package model

import (
	"fmt"

	"golang.org/x/net/context"
)

type UserInterface interface {
	Add(*User) (string, error)
	AddDog(string, string) error
	Get(string) (*User, error)
}

//User holds the information that of the person visiting the dog park
type User struct {
	ID     string   `json:"-"`
	Name   string   `json:"name,omitempty"`
	Email  string   `json:"email,omitempty"`
	ImgURL string   `json:"image,omitempty"`
	Dogs   []string `json:"dogs,omitempty"`
}

type userInterface struct{ UserInterface }

//Add inserts the user into Firebase
func (ui userInterface) Add(u *User) (string, error) {
	ctx := context.Background()
	ref := client.NewRef("users")

	respRef, err := ref.Push(ctx, u)
	if err != nil {
		return "", err
	}
	u.ID = respRef.Key

	return respRef.Key, nil
}

//AddDogKey adds a key associated to Dog already existing in Firebase
//and updates the user.
func (ui userInterface) AddDogKey(userKey, dogKey string) error {
	ctx := context.Background()
	ref := client.NewRef("")

	user, err := ui.Get(userKey)
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

func (ui userInterface) Get(userKey string) (*User, error) {
	ctx := context.Background()
	ref := client.NewRef("users")

	user := new(User)
	err := ref.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
