package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williabk198/shika_app/server/model"
)

type testDogInterface struct{ model.DogInterface }

func (tgi testDogInterface) Add(d *model.Dog) (string, error) {
	return "01234567", nil
}
func (tgi testDogInterface) Get(dogKey string) (*model.Dog, error) {
	return new(model.Dog), nil
}

func TestGetDog(t *testing.T) {

	req := httptest.NewRequest("GET", "/dog/0", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Dog.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}

func TestPosttDog(t *testing.T) {

	req := httptest.NewRequest("Post", "/dog/0", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Dog.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}
