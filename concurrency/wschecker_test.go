package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func wschecker(url string) bool {
	if url != "" {
		return true
	}
	return false
}

func TestCheckWebsite(t *testing.T) {
	urls := []string{
		"http://www.baidu.com",
		"https://github.com",
		"https://www.google.com",
		"",
	}

	expected := map[string]bool{
		"http://www.baidu.com":   true,
		"https://github.com":     true,
		"https://www.google.com": true,
		"":                       false,
	}

	ret := CheckWebsite(wschecker, urls)

	if len(ret) != len(urls) {
		t.Errorf("got %d want %d", len(ret), len(urls))
	}

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("got %v want %v", ret, expected)
	}
}

func sleeper(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "https://this.is.a.test.com"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsite(sleeper, urls)
	}
}
