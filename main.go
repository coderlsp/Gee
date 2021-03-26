package main

import (
	"Gee/gee"
	"log"
	"net/http"
	"time"
)

// middleware example
func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		ctx.Fail(http.StatusInternalServerError, "Internal Server Error.")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	engine := gee.New()
	engine.Use(gee.Logger()) // global middleware
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v2 := engine.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			// expect /hello/hy
			ctx.String(http.StatusOK, "hello %s ,you're at %s\n", ctx.Param("name"), ctx.Path)
		})
	}
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Gee run failed.")
	}
}
