package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
