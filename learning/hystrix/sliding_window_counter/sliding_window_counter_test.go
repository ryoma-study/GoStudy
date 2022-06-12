package sliding_window_counter

import (
	"fmt"
	"testing"
	"time"
)

var snippet = time.Second
var count int32 = 10
var accuracy = time.Second / 10

func take(counter *slidingWindowCounter, count int32) {
	for i := 0; i < int(count); i++ {
		err := counter.Take()
		if err != nil {
			fmt.Println(i, err)
		}
	}
}

func TestSlidingWindowCounter(t *testing.T) {
	counter := New(snippet, count, accuracy)

	for i := 0; i < int(snippet/accuracy)+2; i++ {
		take(counter, count/int32(snippet/accuracy))
		time.Sleep(accuracy)
	}
	//for i := 0; i < int(count); i++ {
	//	err := counter.Take()
	//	if err != nil {
	//		fmt.Println(i, err)
	//	}
	//}
	//time.Sleep(snippet)
	//for i := 0; i < int(count+2); i++ {
	//	err := counter.Take()
	//	if err != nil {
	//		fmt.Println(i, err)
	//	}
	//}
}
