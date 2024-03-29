package router

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jessemolina/gopher/internal/api/mid"
	"github.com/jessemolina/gopher/pkg/web"
)

// Router is a standardize mux for business api.
type Router struct {
	*web.Mux
	mw []web.Middleware
}

// Config defines components required for registering Route Groups.
type Config struct {
	Build    string
	Shutdown chan os.Signal
	Log      *slog.Logger
}

// RouteGroup is defined by instances that can register routes to a web Mux based on Config.
type RouteGroup interface {
	Register(m *web.Mux, cfg Config)
}

// Build composes an http Handler from the given route group based on Config.
func Build(rg RouteGroup, cfg Config) http.Handler {
	// TODO Build a mid.Panics mw and append it to Build mux.
	mux := web.NewMux(
		mid.Logger(cfg.Log),
		mid.Meter(),
		mid.Errors(cfg.Log),
		mid.Signal(cfg.Shutdown),
	)

	rg.Register(mux, cfg)

	return mux
}
