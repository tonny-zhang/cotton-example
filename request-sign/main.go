package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/tonny-zhang/cotton"
)

type reqParam struct {
	name string
	val  string
}
type reqParamList []reqParam

func (s reqParamList) Len() int      { return len(s) }
func (s reqParamList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s reqParamList) Less(i, j int) bool {
	result := strings.Compare(s[i].name, s[j].name)
	if result == 0 {
		result = strings.Compare(s[i].val, s[j].val)
	}
	return result < 0
}

func main() {
	router := cotton.Default()

	router.Use(func(ctx *cotton.Context) {
		ctx.Response.Header().Add("Power by", "cotton")
	})
	{
		g1 := router.Group("/sign")
		g1.Use(func(ctx *cotton.Context) {
			values := ctx.GetAllPostForm()
			if nil != values {
				list := make(reqParamList, 0)
				var sign string
				for k, v := range values {
					if k == "sign" {
						sign = v[0]
					} else {
						list = append(list, reqParam{k, v[0]})
					}
				}

				if sign == "" {
					ctx.JSON(http.StatusOK, cotton.M{
						"code":   4000,
						"errmsg": "not sign",
					})
					ctx.Abort()
					return
				}
				sort.Sort(list)

				strList := make([]string, 0)
				for _, v := range list {
					strList = append(strList, v.name+"="+v.val)
				}
				strSign := strings.Join(strList, "&")
				h := md5.New()
				h.Write([]byte(strSign))
				signGet := hex.EncodeToString(h.Sum(nil))

				fmt.Printf("[sign check] sign = %s, signGet = %s, str = %s\n", sign, signGet, strSign)
				if sign != signGet {
					ctx.JSON(http.StatusOK, cotton.M{
						"code":   4001,
						"errmsg": "sign error",
					})
					ctx.Abort()
				}
			}
		})
		g1.Post("/login", func(ctx *cotton.Context) {
			ctx.String(http.StatusOK, "login")
		})
	}

	router.Run("")
}
