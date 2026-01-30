package main

import (
	"fmt"
	"net/http"
)

type AuroraGeeEngine struct{}

func (engine *AuroraGeeEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func main() {
	engine := &AuroraGeeEngine{}

	err := http.ListenAndServe(":9999", engine)
	if err != nil {
		panic(err)
	}
}
