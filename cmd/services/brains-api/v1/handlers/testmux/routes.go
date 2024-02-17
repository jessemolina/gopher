package testmux

import (
	"github.com/jessemolina/gopher/pkg/web"
)

func Routes(m *web.Mux, mw ...web.Middleware) {

	// logic for registering DDD service/core

	m.GET("/test", HandleTest)
}
