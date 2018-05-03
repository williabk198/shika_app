package model

import (
	"testing"
)

func TestInit(t *testing.T) {
	if ref == nil {
		t.Error("ref was nil. Expected a *db.Ref")
	}
}
