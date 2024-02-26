package testmux

import (
	"os"

	"github.com/jessemolina/gopher/pkg/web"
)

type Config struct {
	shutdown chan os.Signal
}

func Routes(m *web.Mux) {

	// logic for registering DDD service/core

	m.GET("/test", HandleTest)

}
