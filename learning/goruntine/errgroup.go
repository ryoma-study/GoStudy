package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	group, _ := errgroup.WithContext(context.Background())

	for index := 0; index < 3; index++ {
		indexTemp := index
		group.Go(func() error {
			fmt.Printf("goroutine %d done!\n", indexTemp)

			//if indexTemp == 1 {
			//	return errors.New("the index is number 1!")
			//}
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("some goroutine err: ", err)
	} else {
		fmt.Println("all goroutine done success!")
	}
}
