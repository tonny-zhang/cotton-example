package main

import (
	"fmt"
	"net/http"

	"github.com/tonny-zhang/cotton"
)

func main() {
	r := cotton.NewRouter()

	{
		g1 := r.Group("/v1")
		// custom group notfound
		g1.NotFound(func(ctx *cotton.Context) {
			ctx.String(http.StatusNotFound, "group v1 page not found")
		})

		// custom group middleware
		g1.Use(func(ctx *cotton.Context) {
			fmt.Println("g1 middleware")
			ctx.Response.Header().Add("group", "v1")
		})
		g1.Get("/one", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "/v1/one")
		})
	}

	{
		g1 := r.Group("/v2")
		g1.NotFound(func(ctx *cotton.Context) {
			ctx.String(http.StatusNotFound, "group v2 page not found")
		})
		g1.Use(func(ctx *cotton.Context) {
			ctx.Response.Header().Add("group", "v2")
		})
		g1.Get("/one", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "/v2/one")
		})
	}

	r.Run("")
}
