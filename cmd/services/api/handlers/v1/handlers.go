package v1

import (
	"log"
	"net/http"
)


// TODO move to debug mux
func Readiness(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}
	statusCode := http.StatusOK

	if err := response(w, statusCode, data); err != nil{
		log.Println("ERROR ", err)
	}

	log.Println("readiness check")
}

// TODO create server handler

// TODO read server handler

// TODO update server handler

// TODO delete server handler
