package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine https://geektutu.com/post/gee-day1.html
// Engine is the uni handler for all requests
type Engine struct{}

// 在实现 Engine 之前，调用 http.HandleFunc 实现了路由和 Handler 的映射，也就是只能针对具体的路由写处理逻辑
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
