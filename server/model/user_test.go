package model

import (
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func TestAdd(t *testing.T) {
	ref = ref.Child("test")
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
	ref = ref.Child("test")
	testUser := &User{
		Name: "Test User 2",
	}

	userRef, _ := testUser.Add()
	testDog := Dog{
		Name:  "Harley",
		Breed: "Pit Bull Mix",
	}

	err := testUser.AddDogKey(userRef.Key, "3D645BD7FA")
	if err != nil {
		if !strings.HasPrefix(err.Error(), "Error associating") {
			t.Errorf("Expected an error saying the dog was not found.\nGot: %v", err)
		}
	}

	dogRef, _ := testDog.Add()

	err = testUser.AddDogKey(userRef.Key, dogRef.Key)
	if err != nil {
		t.Errorf("Was not expecting an error but got one.")
	}
	ref.Delete(context.Background())
}

func TestUpdateCheckIn(t *testing.T) {
	ref = ref.Child("test")
	testUser := &User{
		Name: "Tester",
	}

	userRef, _ := testUser.Add()
	err := testUser.UpdateCheckIn(userRef, true)
	if err != nil {
		t.Errorf("An error occured where none were expected\nError: %v", err)
	}
	if !testUser.CheckedIn {
		t.Error("Users CheckedIn value was false. Was expecting true")
	}
	ref.Delete(context.Background())
}

func TestSetArea(t *testing.T) {
	ref = ref.Child("test")
	testUser := &User{
		Name: "Tester",
	}

	userRef, _ := testUser.Add()
	err := testUser.SetArea(userRef, "Test")
	if err != nil {
		t.Errorf("An error occured when none was expected\nError: %v", err)
	}
	if testUser.Area != "Test" {
		t.Errorf("Users Area value was %s but expected %s", testUser.Area, "Test")
	}
	ref.Delete(context.Background())
}
