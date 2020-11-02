package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

// SpyStore defines a spy Store
type SpyStore struct {
	response string
	t        *testing.T
}

// SpyResonseWriter defines new ResponseWriter
type SpyResponseWriter struct {
	written bool
}

// rewrite Header()
func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return nil
}

// rewrite Write()
func (w *SpyResponseWriter) Write([]byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

// rewrite WriteHeader()
func (w *SpyResponseWriter) WriteHeader(code int) {
	w.written = true
}

// Fetch fetchs the data from SpyStore
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var ret string
		for _, resp := range s.response {

			select {
			case <-ctx.Done():
				s.t.Log("store cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				ret += string(resp)
			}
		}
		data <- ret
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
