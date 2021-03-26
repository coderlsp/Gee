package gee

import (
	"net/http"
	"strings"
)

// roots key. Example: roots['GET'] roots['POST']
// handlers key. Example: handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
type route struct {
	roots    map[string]*node       //
	handlers map[string]HandlerFunc // route mapping to handler
}

// newRoute is the constructor of gee.route
func newRoute() *route {
	return &route{
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

func (r *route) addRoute(method, pattern string, handler HandlerFunc) {
	//assert(method != "", "HTTP method can not be empty.")
	//assert(pattern[0] == '/', "pattern must begin with '/'.")
	//assert(handler != nil, "handler can not be nil")

	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].addRoute(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// Obtain node and params according to request method and request path.
func (r *route) getRoute(method, path string) (n *node, params map[string]string) {
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
func (r *route) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 Not Found:%s\n", ctx.Path)
		})
	}
	ctx.Next()
}
