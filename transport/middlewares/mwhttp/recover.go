package mwhttp

import (
	"github.com/kkeigenkai/clay/v3/server/middlewares/mwhttp"
)

// Recover recovers HTTP server from handlers' panics.
func Recover(logger interface{}) Middleware {
	return mwhttp.Recover(logger)
}
