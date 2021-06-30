package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tonny-zhang/cotton"
)

func main() {
	router := cotton.Default()

	router.Get("/a", func(ctx *cotton.Context) {
		fmt.Println("req a1")
		time.Sleep(time.Second * 2)

		ctx.String(http.StatusOK, "router a")
		fmt.Println("req a2")
	})
	router.Get("/b", func(ctx *cotton.Context) {
		fmt.Println("req b1")
		time.Sleep(time.Second * 3)
		ctx.String(http.StatusOK, "router b")
		fmt.Println("req b1")
	})
	router.Run("")
}
