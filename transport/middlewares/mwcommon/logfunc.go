package mwcommon

import (
	"context"

	"github.com/kkeigenkai/clay/v3/server/middlewares/mwcommon"
)

func GetLogFunc(logger interface{}) func(context.Context, string) {
	return mwcommon.GetLogFunc(logger)
}
