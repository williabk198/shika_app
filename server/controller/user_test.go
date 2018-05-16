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

	"github.com/williabk198/shika_app/server/db"
	"github.com/williabk198/shika_app/server/model"
	"google.golang.org/appengine/aetest"
)

type testUserImpl struct{ db.UserInterface }

func (tui testUserImpl) Add(u *model.User) (string, error) {
	return "89012345", nil
}
func (tui testUserImpl) Get(key string) (*model.User, error) {
	return new(model.User), nil
}

func TestGetUser(t *testing.T) {

	aeInst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer aeInst.Close()

	req, err := aeInst.NewRequest("GET", "/user/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(User.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Wanted: %d",
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

	testUser := new(model.User)
	err = json.Unmarshal(resp, testUser)
	if err != nil {
		t.Fatalf(
			"Unexpected error unmarshalling response\n%v",
			err,
		)
	}

	expectedUser := model.User{}
	if testUser.ID != expectedUser.ID {
		t.Errorf("User is not the same! Got: %v, Wanted: %v",
			testUser.ID,
			expectedUser.ID,
		)
	}
}

func TestPostUser(t *testing.T) {

	aeInst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer aeInst.Close()

	form := url.Values{
		"name":  {"Brandon Williams"},
		"email": {"bkwilliams92@yahoo.com"},
	}

	formBuffer := new(bytes.Buffer)
	io.WriteString(formBuffer, form.Encode())

	req, err := aeInst.NewRequest("POST", "/user", formBuffer)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(User.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusOK,
		)
	}

	if w.Body.String() != "89012345" {
		t.Errorf(
			"Recieved unexpected value. Got: %s, Wanted: %s",
			w.Body.String(),
			"89012345",
		)
	}
}
