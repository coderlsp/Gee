package gee

import (
	"log"
	"net/http"
	"strings"
)

// roots key. Example: roots['GET'] roots['POST']
// handlers key. Example: handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
type router struct {
	roots    map[string]*node       //
	handlers map[string]HandlerFunc // router mapping to handler
}

// Create a router object.
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one '*' is allowed.
func parsePattern(pattern string) (parts []string) {
	vs := strings.Split(pattern, "/")
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
		}
		if len(item) > 0 && item[0] == '*' {
			break
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	assert(method != "", "HTTP method can not be empty.")
	assert(pattern[0] == '/', "pattern must begin with '/'.")
	assert(handler != nil, "handler can not be nil")

	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].addRoute(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
	log.Printf("Router %4s - %4s\n", method, pattern)
}

// Obtain node and params according to request method and request path.
func (r *router) getRoute(method, path string) (n *node, params map[string]string) {
	params = make(map[string]string)
	searchParts := parsePattern(path)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n = root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// Handle request based on Context.
func (r *router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		r.handlers[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 Not Found:%s\n", ctx.Path)
	}
}
