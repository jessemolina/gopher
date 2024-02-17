package mid

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jessemolina/gopher/pkg/web"
)

// Logger is a middleware function for logging api requests.
func Logger(log *slog.Logger) web.Middleware {

	// Makes use of closure to reference web.Handler from web.wrapMiddleware scope.
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// TODO Determine whether any data from context.Context should be logged.
			// TODO Determine if URL Path query needs to be formatted and logged.

			log.Info("request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Info("request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			return err
		}
		return h

	}
	return m
}
