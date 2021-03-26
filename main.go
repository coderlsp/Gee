package main

import (
	"Gee/gee"
	"log"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := engine.Group("/v1")
	{
		v1.GET("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(ctx *gee.Context) {
			// expect /hello?name=hy
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}
	v2 := engine.Group("/v2")
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			// expect /hello/hy
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
		v2.POST("/login", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Gee run failed.")
	}
}
