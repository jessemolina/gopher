package v1

import (
	"encoding/json"
	"net/http"
)

// convert data stucture to http response with proper header and status code
func response(w http.ResponseWriter, statusCode int, data interface{}) error {

	// convert the response to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// set the content type and headers after successfully marshalling
	w.Header().Set("Content-Type", "application/json")

	// write the status code to the response
	w.WriteHeader(statusCode)

	// send the results back to the client
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
