package model

import (
	"testing"
)

func TestMain(m *testing.M) {
	ref = ref.Child("test")
	m.Run()
}

func TestInit(t *testing.T) {
	if ref == nil {
		t.Error("ref was nil. Expected a *db.Ref")
	}
}
