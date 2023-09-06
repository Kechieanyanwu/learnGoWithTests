package syncing

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Value() int {
	return c.count
}

// Increments counter from one goroutine at a time
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func NewCounter() *Counter {
	return &Counter{}
}
