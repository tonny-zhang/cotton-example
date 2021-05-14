package main

import (
	"net/http"

	"github.com/tonny-zhang/cotton"
)

func main() {
	r := cotton.Default()

	r.NotFound(func(ctx *cotton.Context) {
		ctx.String(http.StatusNotFound, "custom 404 page not found")
	})
	r.Run("")
}
