package sync_benckmark

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func BenchmarkAtomic(t *testing.B) {
	//cfg := &Config{}
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			//cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4}
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load().(*Config)
				fmt.Sprintf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(t *testing.B) {
	var cfg *Config
	var l sync.RWMutex

	go func() {
		i := 0
		for {
			i++
			l.Lock()
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			l.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				l.RLock()
				fmt.Sprintf("%v\n", cfg)
				l.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
