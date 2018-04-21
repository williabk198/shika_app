package model

import (
	"testing"

	"golang.org/x/net/context"
)

func TestInit(t *testing.T) {
	ctx := context.Background()
	ref := dbClient.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		t.Errorf("Error reading from database: %v", err)
	}
}
