package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"google.golang.org/appengine/aetest"

	"github.com/williabk198/shika_app/server/db"
	"github.com/williabk198/shika_app/server/model"
)

type testDogImpl struct{ db.DogInterface }

func (tgi testDogImpl) Add(d *model.Dog) (string, error) {
	return "01234567", nil
}
func (tgi testDogImpl) Get(dogKey string) (*model.Dog, error) {
	return new(model.Dog), nil
}

func TestGetDog(t *testing.T) {

	aeInst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer aeInst.Close()

	req, err := aeInst.NewRequest("GET", "/dog/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Dog.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"Unexpected status code recieved\nGot: %d Wanted: %d",
			w.Code,
			http.StatusOK,
		)
	}

	resp, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf(
			"Unexpected error reading test response\n%v",
			err,
		)
	}

	testDog := new(model.Dog)
	err = json.Unmarshal(resp, testDog)
	if err != nil {
		t.Fatalf(
			"Unexpected error unmarshalling response\n%v",
			err,
		)
	}

	if *testDog != (model.Dog{}) {
		t.Errorf("Dog are not the same! Got: %v, Wanted: %v",
			testDog,
			new(model.Dog),
		)
	}
}

func TestPosttDog(t *testing.T) {

	aeInst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer aeInst.Close()

	form := url.Values{"name": {"Borksly"}, "breed": {"doggo"},
		"sex": {"Male"},
	}

	formBuffer := new(bytes.Buffer)
	io.WriteString(formBuffer, form.Encode())

	req, err := aeInst.NewRequest("POST", "/dog", formBuffer)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "test 12345678")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Dog.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Wanted: %d",
			w.Code,
			http.StatusOK,
		)
	}

	if w.Body.String() != "01234567" {
		t.Errorf(
			"Unexpected response recieve. Got: %s, Wanted: %s",
			w.Body.String(),
			"01234567",
		)
	}
}
