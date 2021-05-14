package main

import (
	"net/http"

	"github.com/tonny-zhang/cotton"
)

func main() {
	r := cotton.Default()

	r.Get("/user/:name", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "user name = "+ctx.Param("name"))
	})
	r.Get("/user/:name/:id", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "user id = "+ctx.Param("id")+" name = "+ctx.Param("name"))
	})
	r.Get("/user/:name/:id/one", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "one user id = "+ctx.Param("id")+" name = "+ctx.Param("name"))
	})
	r.Get("/user/:name/:id/two", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "two user id = "+ctx.Param("id")+" name = "+ctx.Param("name"))
	})
	r.Post("/user/:id", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "hello post "+ctx.Param("id"))
	})

	r.Get("/info/*file", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "info file = "+ctx.Param("file"))
	})
	r.Run("")
}
