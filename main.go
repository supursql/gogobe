package main

import (
	"gogo"
	"net/http"
)

func main() {
	r := gogo.New()
	r.GET("/", func(c *gogo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(c *gogo.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gogo.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s", c.Param("name"), c.Path)
	})

	r.GET("assets/*filepath", func(c *gogo.Context) {
		c.JSON(http.StatusOK, gogo.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")
}
