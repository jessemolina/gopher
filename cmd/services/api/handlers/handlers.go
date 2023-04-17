package handlers

import (
	"net/http"
	"net/http/pprof"
)

// TODO create a debug pprof handlers
func DebugMux() http.Handler{
	mux := http.NewServeMux()

	// register standard debug endpoints
	mux.HandleFunc("/debug/pprof", pprof.Index)

   return mux
}
