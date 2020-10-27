package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Countdown(buf, &CountdownOperationSpy{})

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		ssp := &CountdownOperationSpy{}
		Countdown(ssp, ssp)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, ssp.Calls) {
			t.Errorf("wanted calls %v got %v", want, ssp.Calls)
		}
	})
}
