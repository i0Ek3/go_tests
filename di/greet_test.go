package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "i0Ek3"
	buf := bytes.Buffer{}
	Greet(&buf, name)

	got := buf.String()
	want := "Hello " + name

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
