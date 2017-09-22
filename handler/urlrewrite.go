package handler

import (
	"net/http"

	"github.com/go-ozzo/ozzo-routing"
)

func Urlrewrite(status int) routing.Handler {
	return func(ctx *routing.Context) error {
		if ctx.Request.URL.Path == "/bob" {
			if ctx.Request.Method != "GET" {
				status = http.StatusTemporaryRedirect
			}
			http.Redirect(ctx.Response, ctx.Request, "/mary", status)
			ctx.Abort()
		}
		return nil
	}
}
