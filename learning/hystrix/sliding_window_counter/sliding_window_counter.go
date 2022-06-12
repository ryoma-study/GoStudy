package sliding_window_counter

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	ratelimit_kit "github.com/ulovecode/ratelimit-kit"
)

var (
	once sync.Once
)

var _ ratelimit_kit.RateLimiter = &slidingWindowCounter{}

type slidingWindowCounter struct {
	// 滑动窗口
	accuracy time.Duration
	// 固定窗口
	snippet time.Duration
	// 允许请求数
	allowRequests int32
	// 当前请求数
	currentRequests int32
	// 当前窗口请求数：每隔一个窗口需要移除
	incurRequests int32
	// 利用 buffered chan 记录滑动窗口数据
	durationRequests chan int32
}

func New(snippet time.Duration, allowRequests int32, accuracy time.Duration) *slidingWindowCounter {
	return &slidingWindowCounter{snippet: snippet, allowRequests: allowRequests, accuracy: accuracy, durationRequests: make(chan int32, snippet/accuracy)}
}

func (l *slidingWindowCounter) Take() error {
	once.Do(func() {
		go sliding(l)
		go calculate(l)
	})
	curRequest := atomic.LoadInt32(&l.currentRequests)
	if curRequest >= l.allowRequests {
		return ratelimit_kit.ErrExceededLimit
	}
	// 固定窗口计数
	if !atomic.CompareAndSwapInt32(&l.currentRequests, curRequest, curRequest+1) {
		return ratelimit_kit.ErrExceededLimit
	}
	// 当前窗口计数
	atomic.AddInt32(&l.incurRequests, 1)
	return nil
}

func sliding(l *slidingWindowCounter) {
	for {
		select {
		case <-time.After(l.accuracy):
			// 滑动窗口：移除数据
			t := atomic.SwapInt32(&l.incurRequests, 0)
			//fmt.Printf("reset sliding window counter, incurRequests: %d\n", t)
			l.durationRequests <- t
		}
	}
}

func calculate(l *slidingWindowCounter) {
	for {
		<-time.After(l.accuracy)
		//fmt.Printf("durationRequests waiting %d %d\n", len(l.durationRequests), cap(l.durationRequests))
		if len(l.durationRequests) == cap(l.durationRequests) {
			break
		}
	}
	fmt.Printf("durationRequests enough %d\n", cap(l.durationRequests))
	for {
		<-time.After(l.accuracy)
		t := <-l.durationRequests
		if t != 0 {
			fmt.Println("durationRequests start consume")
			atomic.AddInt32(&l.currentRequests, -t)
		}
	}
}
