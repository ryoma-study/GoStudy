package fixed_window_counter

import (
	"fmt"
	"testing"
	"time"
)

var count int32 = 10
var snippet = time.Second

func TestFixedWindowCounter(t *testing.T) {
	counter := New(snippet, count)

	for i := 0; i < int(count+2); i++ {
		err := counter.Take()
		if err != nil {
			fmt.Println(i, err)
		}
	}

	time.Sleep(snippet * 2)

	for i := 0; i < int(count+2); i++ {
		err := counter.Take()
		if err != nil {
			fmt.Println(i, err)
		}
	}
}
