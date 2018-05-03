package model

import (
	"testing"

	"golang.org/x/net/context"
)

func TestDogAdd(t *testing.T) {
	ref = ref.Child("test")
	testDog := &Dog{
		Name: "Test Dog 1",
	}

	_, err := testDog.Add()
	if err != nil {
		t.Errorf("Error occured but none was expected\n%v", err)
	}
	ref.Delete(context.Background())
}

func TestDogUpdateCheckIn(t *testing.T) {
	ref = ref.Child("test")
	testDog := &Dog{
		Name: "Tester",
	}

	dogRef, _ := testDog.Add()
	err := testDog.UpdateCheckIn(dogRef, true)
	if err != nil {
		t.Errorf("An error occured where none were expected\nError: %v", err)
	}
	if !testDog.CheckedIn {
		t.Error("Dogs CheckedIn value was false. Was expecting true")
	}
	ref.Delete(context.Background())
}
