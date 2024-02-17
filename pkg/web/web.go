package web

import (
	"os"
	"syscall"
)

// App is a custom web application.
type App struct {
	Mux *Mux
	shutdown chan os.Signal
}

// NewApp is creates a new custom web application.
func NewApp(m *Mux) *App {
	app := &App{
		Mux: m,
	}

	return app
}

// SignalShutdown sends app a terminal signal for a graceful shutdown.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
