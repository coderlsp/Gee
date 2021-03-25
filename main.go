package main

import (
	"Gee/gee"
	"log"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	engine.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"filepath": ctx.Param("filepath"),
		})
	})

	engine.GET("/hello/:name/b", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"name":  ctx.Param("name"),
			"place": "b",
		})
	})
	engine.GET("/hello/:name/c", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"name":  ctx.Param("name"),
			"place": "c",
		})
	})
	engine.GET("/hello/:you/b", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"you": ctx.Param("you"),
		})
	})
	engine.GET("/hello", func(ctx *gee.Context) {
		// expect /hello?name=hy
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Gee run failed.")
	}
}
