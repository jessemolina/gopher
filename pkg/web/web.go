package web

import (
	"context"
	"fmt"
	"net/http"
)

// TODO Embed web App with shutdown signal.
// TODO Consider whether to use custom Ctx with embeded ResponseWriter for response flex (i.e. JSON vs XML).

// Handler is a handle function signature that accounts for context and error.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is a custom web application.
type App struct {
	*http.ServeMux
	mw []Middleware
}

// NewApp is creates a new custom web application.
func NewApp(mw ...Middleware) *App {
	app := &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}

	return app
}

// handle wraps middleware functions to the Handler and registers it to the serve mux.
func (a *App) handle(method string, path string, handler Handler) {
	handler = wrapMiddleware(handler, a.mw)

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

	a.HandleFunc(path, h)

}

// GET registers a GET method request to the provided path for the provided custom Handler.
func (a *App) GET(path string, handler Handler) {
	a.handle(http.MethodGet, path, handler)
}
