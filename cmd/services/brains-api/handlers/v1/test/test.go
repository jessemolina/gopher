package test

import (
	"encoding/json"
	"net/http"
)

// HandleTest handles responds status in JSON format.
func HandleTest(w http.ResponseWriter, r *http.Request) {

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
