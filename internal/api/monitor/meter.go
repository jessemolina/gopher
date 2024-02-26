package mid

import (
	"context"
	"net/http"

	"github.com/jessemolina/gopher/internal/api/metrics"
	"github.com/jessemolina/gopher/pkg/web"
)

// Meter is a middleware for measuring api metrics.
func Meter() web.Middleware {

	mw := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			am := metrics.ApiMetrics()
			am.AddRequests(ctx, 1)

			err := handler(ctx, w, r)
			if err != nil {
				am.AddErrors(ctx, 1)
			}

			return err
		}
		return h

	}
	return mw
}
