package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// H 使得JSON格式更加简洁
type H map[string]interface{}

type Context struct {
	// origin object
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	_, err := ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		log.Fatal("write string failed.")
	}
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	_, err := ctx.Writer.Write(data)
	if err != nil {
		log.Fatal("write data failed.")
	}
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	_, err := ctx.Writer.Write([]byte(html))
	if err != nil {
		log.Fatal("write HTML failed.")
	}
}

func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]
	return value
}
