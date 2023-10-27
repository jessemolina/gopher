package handlers

import (
	"log/slog"

	"github.com/jessemolina/gopher/cmd/services/brains-api/handlers/v1/test"
	"github.com/jessemolina/gopher/internal/api/v1/mid"
	"github.com/jessemolina/gopher/pkg/web"
)

type APIMuxConfig struct {
	Log *slog.Logger
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(mid.Logger(cfg.Log))

	app.GET("/test", test.HandleTest)

	return app
}
