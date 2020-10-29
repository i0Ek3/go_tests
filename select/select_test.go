package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("get response", func(t *testing.T) {
		fastOne := mockHTTPTest(8 * time.Millisecond)
		slowOne := mockHTTPTest(12 * time.Millisecond)

		defer fastOne.Close()
		defer slowOne.Close()

		_, err := Racer(fastOne.URL, slowOne.URL)

		assertTimeout(t, err, nil)
	})

	t.Run("get timeout", func(t *testing.T) {
		runTime := mockHTTPTest(5 * time.Second)
		defer runTime.Close()

		_, err := racer(runTime.URL, runTime.URL, 3*time.Second)

		if err == nil {
			t.Errorf("we don't have an error, but we want it")
		}
	})
}

func mockHTTPTest(delay time.Duration) *httptest.Server {
	mockURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return mockURL
}

func assertTimeout(t *testing.T, got, want error) {
	if got != nil {
		t.Fatalf("get error %s but we don't want it", got)
	}

	if got != want {
		t.Errorf("something wrong here, looks timeout")
	}
}

func BenchmarkMockA(b *testing.B) {
	fastOne := mockHTTPTest(8 * time.Millisecond)
	defer fastOne.Close()

	Racer(fastOne.URL, fastOne.URL)
}

func BenchmarkMockB(b *testing.B) {
	fastOne := mockHTTPTest(8 * time.Millisecond)
	defer fastOne.Close()

	racer(fastOne.URL, fastOne.URL, delayTimeout)
}
