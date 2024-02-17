package mid

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/jessemolina/gopher/pkg/web"
)

// Logger is a middleware function for logging api requests.
func Signal(shutdown chan os.Signal) web.Middleware {

	// Makes use of closure to reference web.Handler from web.wrapMiddleware scope.
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			err := handler(ctx, w, r)

			if web.IsShutdown(err) {
				shutdown <- syscall.SIGTERM
			}

			return err
		}
		return h

	}
	return m
}
