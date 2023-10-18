package test

import (
	"net/http"
)

// HandleTest handles a test request.
func HandleTest(w http.ResponseWriter, r *http.Request) {

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	w.Write([]byte(status.Status))

}
