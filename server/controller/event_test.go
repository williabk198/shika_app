package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williabk198/shika_app/server/model"
)

type testEventInterface struct{ model.EventInterface }

func (tei testEventInterface) Add(e *model.Event) (string, error) {
	return "45678901", nil
}

func TestGetEvent(t *testing.T) {
	req := httptest.NewRequest("GET", "/event/0", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Event.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}

func TestPostEvent(t *testing.T) {
	req := httptest.NewRequest("POST", "/event", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Event.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}
