package web

import (
	"context"
	"fmt"
	"net/http"
)

type Mux struct {
	*http.ServeMux
	mw []Middleware
}

// Handler is a handle function signature that accounts for context and error.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func NewMux(mw ...Middleware) *Mux {
	router := &Mux{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}

	return router
}

// handle wraps middleware functions to the Handler and registers it to the serve mux.
func (m *Mux) Handle(method string, path string, handler Handler, mw ...Middleware) {
	handler = Wrap(handler, mw)
	handler = Wrap(handler, m.mw)

	h := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := handler(r.Context(), w, r); err != nil {
			fmt.Println(err)
			return
		}
	}

	m.HandleFunc(path, h)

}

// GET registers a GET method request to the provided path for the provided custom Handler.
func (m *Mux) GET(path string, handler Handler, mw ...Middleware) {
	m.Handle(http.MethodGet, path, handler, mw...)
}
