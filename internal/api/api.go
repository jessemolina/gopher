package api

import (
	"log/slog"
	"os"

	"github.com/jessemolina/gopher/internal/api/monitor"
	"github.com/jessemolina/gopher/internal/api/safeguard"
	"github.com/jessemolina/gopher/pkg/web"
)

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
func Build(rg RouteGroup, cfg Config) *web.Mux {
	// TODO Build a mid.Panics mw and append it to Build mux.
	mux := web.NewMux(
		monitor.Logger(cfg.Log),
		monitor.Meter(),
		safeguard.Errors(cfg.Log),
		safeguard.Signal(cfg.Shutdown),
	)

	rg.Register(mux, cfg)

	return mux
}
