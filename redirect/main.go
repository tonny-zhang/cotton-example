package main

import "github.com/tonny-zhang/cotton"

func main() {
	r := cotton.Default()

	r.Get("/redirect", func(ctx *cotton.Context) {
		urlto := ctx.GetDefaultQuery("url", "https://www.baidu.com")
		ctx.Redirect(302, urlto)
	})
	r.Run("")
}
