package test

import (
	"context"
	"encoding/json"
	"net/http"
)

// HandleTest handles responds status in JSON format.
func HandleTest(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)

	return nil
}
