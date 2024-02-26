package tests

import (
	"github.com/jessemolina/gopher/cmd/services/brains-api/v1/handlers/testmux"
	"github.com/jessemolina/gopher/internal/api"
	"github.com/jessemolina/gopher/pkg/web"
)

func Routes() group {
	return group{}

}

// group implements Router RouteGroup
type group struct{}

// Register adds mux routes to the router.
func (g group) Register(m *web.Mux, cfg api.Config) {
	testmux.Routes(m)
}
