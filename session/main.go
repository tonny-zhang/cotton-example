package main

import (
	"fmt"
	"net/http"

	"github.com/tonny-zhang/cotton"
	session "github.com/tonny-zhang/cotton-session"
)

func main() {
	router := cotton.NewRouter()
	router.Use(cotton.RecoverWithWriter(nil, func(ctx *cotton.Context, err interface{}) {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("[ERROR] %v", err))
	}))
	router.Use(cotton.Logger())
	// router.Use(session.Middleware(""))
	g := router.Group("/session")
	g.Use(session.Middleware(""))
	g.Post("/set", func(ctx *cotton.Context) {
		session.GetSession(ctx).Set("user", "tonny")
		ctx.String(http.StatusOK, "writed")
	})
	g.Get("/get", func(ctx *cotton.Context) {
		v, e := session.GetSession(ctx).Get("user")
		if e == nil {
			ctx.String(http.StatusOK, v.(string))
		} else {
			ctx.String(http.StatusOK, "not found")
		}
	})
	g.Get("/delete", func(ctx *cotton.Context) {
		session.GetSession(ctx).Del("user")

		ctx.String(http.StatusOK, "delete")
	})

	router.Get("/nosession", func(ctx *cotton.Context) {
		v, e := session.GetSession(ctx).Get("user")
		if e == nil {
			ctx.String(http.StatusOK, v.(string))
		} else {
			ctx.String(http.StatusOK, "not found")
		}
	})
	router.Run("")
}
