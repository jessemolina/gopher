package safeguard

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jessemolina/gopher/pkg/web"
)

// Logger is a middleware function for logging api requests.
func Errors(log *slog.Logger) web.Middleware {

	// Makes use of closure to reference web.Handler from web.wrapMiddleware scope.
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			if err := handler(ctx, w, r); err != nil {
				log.Error("request error", "msg", err)

				if web.IsShutdown(err) {
					return err
				}
			}

			return nil
		}
		return h

	}
	return m
}
