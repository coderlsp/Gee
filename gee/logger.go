package gee

import (
	"log"
	"time"
)

// middleware example
func Logger() HandlerFunc {
	return func(ctx *Context) {
		// Start timer
		t := time.Now()
		// Process request
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
