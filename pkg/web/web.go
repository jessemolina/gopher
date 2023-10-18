package web

import "net/http"

// TODO Define web Handler function signature; http request with context.

// TODO Build Mux struct that embeds a mux, shutdown sig, and mw functions.

// TODO Factory Function for creating NewMux

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	app := &App{
		http.NewServeMux(),
	}
	return app
}
