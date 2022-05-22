package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(process("hello GO!"))
}

type result struct {
	record string
	err    error
}

func process(term string) error {
	ctx, _ := context.WithCancel(context.Background())
	ch := make(chan result, 1)
	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()
	//cancel()
	select {
	case <-ctx.Done():
		return errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received: ", result.record)
		return nil
	}
}

//func process(term string) error {
//	record, err := search(term)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println("Received: ", record)
//	return nil
//}

func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}
