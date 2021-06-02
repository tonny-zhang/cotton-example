package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/tonny-zhang/cotton"
)

type timeWriter struct {
	filepathFormat string
	filepath       string
	file           *os.File
}

func (writer *timeWriter) Write(p []byte) (n int, err error) {
	logfilePath := time.Now().Format(writer.filepathFormat)
	if writer.filepath != logfilePath && writer.file != nil {
		writer.file.Close()
		writer.file = nil
		writer.filepath = ""
	}
	if writer.file == nil {
		e := os.MkdirAll(path.Dir(logfilePath), os.ModePerm)
		if e != nil {
			fmt.Println(e)
		}
		writer.filepath = logfilePath
		writer.file, e = os.OpenFile(logfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if e != nil {
			fmt.Println(e)
		}
	}

	return writer.file.Write(p)
}
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	logDir := path.Join(path.Dir(filename), "log")
	writer := &timeWriter{
		filepathFormat: path.Join(logDir, "access.200601021504.log"),
	}

	writerError := &timeWriter{
		filepathFormat: path.Join(logDir, "error.200601021504.log"),
	}

	router := cotton.NewRouter()
	router.Use(cotton.RecoverWithWriter(writerError), cotton.LoggerWidthConf(cotton.LoggerConf{
		Writer: writer,
	}))
	router.Get("/test", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "test")
	})
	router.Get("/panic", func(ctx *cotton.Context) {
		panic("123")
	})

	d1 := router.Domain("a.pilin.com")
	d1.Use(cotton.LoggerWidthConf(cotton.LoggerConf{
		Writer: &timeWriter{
			filepathFormat: path.Join(logDir, "a.pilin.com-200601021504.log"),
		},
	}))
	d1.Get("/test", func(ctx *cotton.Context) {
		ctx.String(http.StatusOK, "d1 test")
	})

	router.Run("")
}
