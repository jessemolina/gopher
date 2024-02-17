package testmux

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/jessemolina/gopher/pkg/web"
)

// HandleTest handles responds status in JSON format.
func HandleTest(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if rand.Float32() < 0.25 {
		return errors.New("Simulated error in test handler")
	}

	status := struct {
		Status string
	}{
		Status: "OKIE DOKIE",
	}

	web.Response(ctx, w, http.StatusOK, status)

	return nil
}
