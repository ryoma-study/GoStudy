package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo`
		c <- `bar`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		println(`Message: ` + <-c)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}
