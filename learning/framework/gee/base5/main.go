package main

import (
	"net/http"

	"gee"
)

// https://geektutu.com/post/gee-day3.html
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/ryoma", func(c *gee.Context) {
		// expect /hello/ryoma
		c.String(http.StatusOK, "hello ryoma, you're at %s\n", c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

//$ curl "http://localhost:9999/hello/geektutu"
//hello geektutu, you're at /hello/geektutu
//
//$ curl "http://localhost:9999/assets/css/geektutu.css"
//{"filepath":"css/geektutu.css"}
