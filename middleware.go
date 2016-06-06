package osgraphql

import (
	"strings"

	"github.com/Unknwon/macaron"
)

// CORS sort out the CORS headers
func CORS() macaron.Handler {
	return func(ctx *macaron.Context) {

		headers := ctx.Resp.Header()

		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", strings.Join([]string{
			"POST",
			"GET",
			"OPTION",
		}, ", "))

		headers.Set("Access-Control-Allow-Headers", strings.Join([]string{
			"authorization",
			"Accept",
			"Content-Type",
			"Content-Length",
		}, ", "))

		ctx.Next()
	}
}
