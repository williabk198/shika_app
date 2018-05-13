package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williabk198/shika_app/server/model"
)

type testVisitorInterface struct{ model.VisitorInterface }

func (tvi testVisitorInterface) Add(v *model.Visitor) (string, error) {
	return "67890123", nil
}
func (tvi testVisitorInterface) Remove(s string) error {
	return nil
}

func TestGetVisitor(t *testing.T) {
	req := httptest.NewRequest("GET", "/visitor/0", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Visitor.Get)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}

func TestPostVisitor(t *testing.T) {
	req := httptest.NewRequest("POST", "/visitor", nil)
	req.Header.Set("Authorization", "Test 12345678")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Visitor.Post)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotImplemented {
		t.Errorf(
			"Unexpected error code recieved\nGot: %d Revieved: %d",
			w.Code,
			http.StatusNotImplemented,
		)
	}
}
