package main

import (
	"sync"
)

// Counter defines a counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter news a pointer type Counter
// Don't initilize it directly in your code
func NewCounter() *Counter {
	return &Counter{}
}

// Incr increases value
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the value
func (c *Counter) Value() int {
	return c.value
}
