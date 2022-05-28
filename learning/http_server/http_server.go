package main

import (
	"fmt"
	"log"
	"net/http"
)

// 通过 http 底层的 server 启动、关闭
func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello World!")
	//})
	//
	//http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello server!")
	})

	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server start error: ", err)
	}
}
