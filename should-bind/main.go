package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/tonny-zhang/cotton"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type bindPerson struct{}

func decode(b []byte, obj interface{}) (err error) {
	arr := strings.Split(string(b), "|")
	if len(arr) == 2 {
		name := arr[0]
		age, e := strconv.Atoi(arr[1])
		if e == nil {
			v := reflect.ValueOf(obj).Elem()
			v.FieldByName("Name").SetString(name)
			v.FieldByName("Age").SetInt(int64(age))
		} else {
			err = e
		}
	} else {
		err = fmt.Errorf("bad request")
	}
	return
}
func (bindPerson) Bind(req *http.Request, obj interface{}) (err error) {
	if req == nil || req.Body == nil {
		err = fmt.Errorf("bad request")
	} else {
		var b []byte
		b, err = ioutil.ReadAll(req.Body)
		if err == nil {
			err = decode(b, obj)
		}
	}
	return
}
func (bindPerson) BindBody(b []byte, obj interface{}) error {
	return decode(b, obj)
}
func main() {
	router := cotton.Default()

	router.Post("/json", func(ctx *cotton.Context) {
		var p person
		e := ctx.ShouldBindWithJSON(&p)
		if nil == e {
			ctx.JSON(http.StatusOK, cotton.M{
				"name": p.Name,
				"age":  p.Age,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, cotton.M{
				"msg": e.Error(),
			})
		}
	})
	router.Post("/json2", func(ctx *cotton.Context) {
		var p person
		e := ctx.ShouldBindWith(&p, bindPerson{})
		if nil == e {
			ctx.JSON(http.StatusOK, cotton.M{
				"name": p.Name,
				"age":  p.Age,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, cotton.M{
				"msg": e.Error(),
			})
		}
	})

	router.Post("/json3", func(ctx *cotton.Context) {
		var p person
		var p1 person
		e := ctx.ShouldBindBodyWithJSON(&p)
		e1 := ctx.ShouldBindBodyWith(&p1, bindPerson{})
		data := cotton.M{
			"p":  p,
			"e":  e,
			"p1": p1,
			"e1": e1,
		}
		ctx.JSON(http.StatusOK, data)
	})
	router.Run("")
}
