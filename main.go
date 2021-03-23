package main

import (
	"Gee/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {

		}
	})
	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	err := engine.Run(":8080")
	if err != nil {
		log.Printf("Addr Error!")
	}
}
