package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %4s\n", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 Not Found:%s\n", ctx.Path)
	}
}
