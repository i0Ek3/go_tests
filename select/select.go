package main

import (
	"errors"
	"net/http"
	"time"
)

var (
	Timeout = errors.New("timeout")

	delayTimeout = time.Duration(10 * time.Millisecond)
)

//// a common version
//// Racer picks fast url
//func Racer(a, b string) string {
//	ea := mockHTTP(a)
//	eb := mockHTTP(b)
//
//	if ea > eb {
//		return a
//	}
//	return b
//}
//
//func mockHTTP(u string) time.Duration {
//	tu := time.Now()
//	http.Get(u)
//	ea := time.Since(tu)
//	return ea
//}

// a select version
// Racer picks fast url
func Racer(a, b string) (string, error) {
	return racer(a, b, delayTimeout)
}

func racer(a, b string, delay time.Duration) (string, error) {
	select {
	case <-mockHTTP(a):
		return a, nil
	case <-mockHTTP(b):
		return b, nil
	case <-time.After(delay):
		return "", Timeout
	}
}

func mockHTTP(u string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(u)
		ch <- true
	}()
	return ch
}
