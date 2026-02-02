package main

import (
	"net/http"

	auroragee "aurora-gee"
)

func main() {
	r := auroragee.New()
	r.GET("/", func(c *auroragee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(c *auroragee.Context) {
		c.JSON(http.StatusOK, auroragee.H{
			"value": "hello world",
		})
	})

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
