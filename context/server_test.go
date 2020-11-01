package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return nil
}

func (w *SpyResponseWriter) Write([]byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(code int) {
	w.written = true
}

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

func TestServer(t *testing.T) {
	data := "hello context"

	t.Run("fetch data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		ser := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)
		response := &SpyResponseWriter{}

		ser.ServeHTTP(response, request)

		if response.written {
			t.Error("response not allow to written")
		}
	})

	t.Run("not allowed", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		ser := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		ser.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}
	})
}
