package model

import (
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func TestAdd(t *testing.T) {
	testUser := &User{
		Name: "Test User 1",
	}
	_, err := testUser.Add()
	if err != nil {
		t.Errorf("Error occured but none was expected\n%v", err)
	}
	ref.Delete(context.Background())
}

func TestAddDog(t *testing.T) {
	testUser := &User{
		Name: "Test User 2",
	}

	testUser.Add()
	testDog := Dog{
		Name:  "Harley",
		Breed: "Pit Bull Mix",
	}

	err := testUser.AddDogKey("3D645BD7FA")
	if err != nil {
		if !strings.HasPrefix(err.Error(), "Error associating") {
			t.Errorf("Expected an error saying the dog was not found.\nGot: %v", err)
		}
	}

	dogRef, _ := testDog.Add()

	err = testUser.AddDogKey(dogRef.Key)
	if err != nil {
		t.Errorf("Was not expecting an error but got one.")
	}
	ref.Delete(context.Background())
}
