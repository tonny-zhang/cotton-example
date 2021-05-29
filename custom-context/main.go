package main

import (
	"net/http"

	"github.com/tonny-zhang/cotton"
)

// ContextCustom custom context
type ContextCustom struct {
	*cotton.Context
}

func (ctx ContextCustom) getFullPath() string {
	return ctx.Request.URL.Path
}
func upgradeContext(ctx *cotton.Context) ContextCustom {
	return ContextCustom{
		Context: ctx,
	}
}
func main() {
	router := cotton.Default()

	router.Get("/*test", func(ctx *cotton.Context) {
		ctxCustom := upgradeContext(ctx)

		ctxCustom.String(http.StatusOK, "fullpath = "+ctxCustom.getFullPath())
	})
	router.Run("")
}
