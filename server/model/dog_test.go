package model

import (
	"testing"

	"golang.org/x/net/context"
)

func TestDogAdd(t *testing.T) {
	testDog := &Dog{
		Name: "Test Dog 1",
	}

	_, err := testDog.Add()
	if err != nil {
		t.Errorf("Error occured but none was expected\n%v", err)
	}
	ref.Delete(context.Background())
}
