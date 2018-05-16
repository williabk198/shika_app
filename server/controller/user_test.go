package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williabk198/shika_app/server/db"
	"github.com/williabk198/shika_app/server/model"
)

type testUserImpl struct{ db.UserInterface }

func (tui testUserImpl) Add(u *model.User) (string, error) {
	return "89012345", nil
}
func (tui testUserImpl) Get(key string) (*model.User, error) {
	return new(model.User), nil
}

func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/user/0", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(User.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}

func TestPostUser(t *testing.T) {
	req := httptest.NewRequest("POST", "/user", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(User.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}
