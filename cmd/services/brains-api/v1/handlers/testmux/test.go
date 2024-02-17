package testmux

import (
	"context"
	"net/http"

	"github.com/jessemolina/gopher/pkg/web"
)

// HandleTest handles responds status in JSON format.
func HandleTest(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OKIE DOKIE",
	}

	web.Response(ctx, w, http.StatusOK, status)

	return nil
}
