package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	want := 1000

	var wg sync.WaitGroup
	wg.Add(want)

	t.Run("concurrent safely", func(t *testing.T) {
		t.Helper()
		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				c.Incr()
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t, c, want)
	})

}

func assertCounter(t *testing.T, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d but we want %d", got.Value(), want)
	}

}
