package limit_algorithm

import (
	"sync"
	"time"
)

var mu sync.Mutex

type Counter struct {
	rate  int
	begin time.Time
	cycle time.Duration
	count int
}

func (c *Counter) Reset() {
	c.begin = time.Now()
	c.count = 0
}

func (c *Counter) Allow() bool {
	mu.Lock()
	defer mu.Unlock()

	if c.count+1 > c.rate {
		if time.Now().Sub(c.begin) > c.cycle {
			c.Reset()
			return true
		} else {
			return false
		}
	} else {
		c.count++
		return true
	}
}

func (c *Counter) Set(r int, cycle time.Duration) {
	c.cycle = cycle
	c.rate = r
	c.begin = time.Now()
	c.count = 0
}
