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

	// use memory session
	g := router.Group("/session")
	mgrMemory := session.NewMemoryMgr()
	g.Use(session.Middleware(mgrMemory))

	g.Post("/set", func(ctx *cotton.Context) {
		ss := session.GetSession(ctx)
		ss.Set("user", ctx.GetPostForm("user"))
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

	// use redis session
	mgrRedis := session.NewRedisMgr("localhost:6379", 0, "")
	gRedis := router.Group("/redis")
	gRedis.Use(session.Middleware(mgrRedis))
	gRedis.Post("/set", func(ctx *cotton.Context) {
		ss := session.GetSession(ctx)
		ss.Set("user", ctx.GetPostForm("user"))
		ss.Save()
		ctx.String(http.StatusOK, "writed")
	})
	gRedis.Get("/get", func(ctx *cotton.Context) {
		v, e := session.GetSession(ctx).Get("user")
		if e == nil {
			ctx.String(http.StatusOK, v.(string))
		} else {
			ctx.String(http.StatusOK, "not found")
		}
	})
	gRedis.Get("/delete", func(ctx *cotton.Context) {
		ss := session.GetSession(ctx)
		ss.Del("user")
		ss.Save()

		ctx.String(http.StatusOK, "delete")
	})

	// no use session with panic
	router.Get("/nosession", func(ctx *cotton.Context) {
		v, e := session.GetSession(ctx).Get("user")
		if e == nil {
			ctx.String(http.StatusOK, v.(string))
		} else {
			ctx.String(http.StatusOK, "not found")
		}
	})
	// no use session with not panic
	router.Get("/nosession/nopanic", func(ctx *cotton.Context) {
		if session.HasUsedSession(ctx) {
			ss := session.GetSession(ctx)
			v, e := ss.Get("user")
			if e == nil {
				ctx.String(http.StatusOK, v.(string))
			} else {
				ctx.String(http.StatusOK, "not found")
			}
		} else {
			ctx.String(http.StatusOK, "not use session")
		}
	})
	router.Run("")
}
