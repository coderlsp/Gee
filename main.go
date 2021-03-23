package main

import (
	"Gee/gee"
	"log"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee.</h1>")
	})
	engine.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s,you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	engine.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Printf("Addr Error!")
	}
}
