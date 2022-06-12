package fixed_window_counter

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

var _ ratelimit_kit.RateLimiter = &fixedWindowCounter{}

type fixedWindowCounter struct {
	// 固定窗口
	snippet time.Duration
	// 允许请求数
	allowRequests int32
	// 当前请求数
	currentRequests int32
}

func New(snippet time.Duration, allowRequests int32) *fixedWindowCounter {
	return &fixedWindowCounter{snippet: snippet, allowRequests: allowRequests}
}

func (l *fixedWindowCounter) Take() error {
	// 固定时间重置即可
	once.Do(func() {
		go func() {
			for {
				select {
				case <-time.After(l.snippet):
					fmt.Println("reset fixed window counter")
					atomic.StoreInt32(&l.currentRequests, 0)
				}
			}
		}()
	})

	// 判断当前窗口内是否正常
	curRequest := atomic.LoadInt32(&l.currentRequests)
	if curRequest >= l.allowRequests {
		return ratelimit_kit.ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&l.currentRequests, curRequest, curRequest+1) {
		return ratelimit_kit.ErrExceededLimit
	}
	return nil
}
