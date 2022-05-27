package main

import (
	"sync"
)

func main() {

}

type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

func Poller1(res *Resources) {
	for {
		// get the least recently-polled Resource
		// and mark it as being polled
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil {
			continue
		}

		// poll the URL

		// update the Resource's polling and lastPolled
		res.lock.Lock()
		r.polling = false
		//r.lastPolled = time.Nanoseconds()
		res.lock.Unlock()
	}
}

type Resource1 string

func Poller2(in, out chan *Resource1) {
	for r := range in {
		// poll the URL

		// send the processed Resource to out
		out <- r
	}
}
