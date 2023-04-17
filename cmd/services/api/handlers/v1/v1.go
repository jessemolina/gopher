package v1

import "net/http"

func Routes(mux *http.ServeMux) {
   // TODO move to debug
   mux.HandleFunc("/debug/ready", Readiness)
}
