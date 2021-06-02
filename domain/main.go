package main

import (
	"fmt"
	"net/http"

	"github.com/tonny-zhang/cotton"
)

func main() {
	router := cotton.NewRouter()
	router.Use(cotton.Recover())
	{
		d1 := router.Domain("www.pilin.com", func(ctx *cotton.Context) {
			fmt.Println("middleware d1", ctx.Request.Host)
		}, cotton.LoggerWidthConf(cotton.LoggerConf{
			Formatter: func(param cotton.LoggerFormatterParam, ctx *cotton.Context) string {
				return fmt.Sprintf("[www-INFO] %v\t%13s %6s %3d %10v %s \n",
					param.TimeStamp.Format("2006/01/02 15:04:05"),
					param.ClientIP,
					param.Method,
					param.StatusCode,
					param.Latency,
					param.Path,
				)
			},
		}))
		d1.NotFound(func(ctx *cotton.Context) {
			ctx.String(http.StatusNotFound, "d1 page not found")
		})
		d1.Get("/test", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "www test")
		})
		g1 := d1.Group("/v1")
		g1.NotFound(func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "d1 v1 page not found")
		})
		g1.Get("/test", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "www v1 test")
		})
	}
	{
		d2 := router.Domain("a.pilin.com", cotton.RecoverWithWriter(nil, func(ctx *cotton.Context, err interface{}) {
			fmt.Println("[a-recover]", err)
		}), cotton.LoggerWidthConf(cotton.LoggerConf{
			Formatter: func(param cotton.LoggerFormatterParam, ctx *cotton.Context) string {
				return fmt.Sprintf("[a-INFO] %v\t%13s %6s %3d %10v %s \n",
					param.TimeStamp.Format("2006/01/02 15:04:05"),
					param.ClientIP,
					param.Method,
					param.StatusCode,
					param.Latency,
					param.Path,
				)
			},
		}))
		d2.NotFound(func(ctx *cotton.Context) {
			ctx.String(http.StatusNotFound, "d2 page not found")
		})
		d2.Get("/test", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "a test")
		})
		d2.Get("/panic", func(ctx *cotton.Context) {
			panic("test")
		})
	}
	router.Get("/test", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "test")
	})
	router.Run("")
}
