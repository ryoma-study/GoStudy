package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}
