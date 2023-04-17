package handlers

import (
	"net/http"
	"net/http/pprof"

	"github.com/jessemolina/gopher/cmd/services/api/handlers/v1"
)

// create a debug pprof handlers
func DebugMux() http.Handler {
	mux := http.NewServeMux()

	// register standard debug endpoints
	mux.HandleFunc("/debug/pprof", pprof.Index)

	return mux
}

// TODO create an api handlers
func APIMux() http.Handler {
	mux := http.NewServeMux()

	// TODO api crud handlers
	v1.Routes(mux)

	return mux
}
