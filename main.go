package main

import (
	"fmt"
	"net/http"

	auroragee "aurora-gee"
)

func main() {
	r := auroragee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
