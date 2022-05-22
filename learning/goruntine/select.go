package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, GopherCon SG")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "hello, GopherCon SG")
//	})
//
//	go func() {
//		if err := http.ListenAndServe(":8080", nil); err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	// 阻塞，避免主进程结束
//	select {}
//}
