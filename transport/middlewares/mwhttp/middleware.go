package mwhttp

import (
	"github.com/kkeigenkai/clay/v3/server/middlewares/mwhttp"
)

// Middleware is the HTTP middleware type.
// It processes the request (potentially mutating it) and
// gives control to the underlying handler.
type Middleware = mwhttp.Middleware
