package web

import ()

// Middleware is function signature for the custom web middleware type.
type Middleware func(Handler) Handler

// wrapMiddleware wraps a handler function with each middlware function.
// The mw functions are wrapped in reverse order; mw funcs run from left to right.
func Wrap(handler Handler, mw []Middleware) Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}
	return handler
}
