package main

import (
	"gogo"
	"net/http"
)

func main() {
	r := gogo.Default()
	r.GET("/", func(c *gogo.Context) {
		c.String(http.StatusOK, "Hello gogo\n")
	})

	r.GET("/panic", func(c *gogo.Context) {
		names := []string{"gogo"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9000")
}
