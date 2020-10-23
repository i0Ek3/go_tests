package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("Espanol", func(t *testing.T) {
		got := Hello("i0Ek3", "Espanol")
		want := "Hola, i0Ek3"

		assertMsg(t, got, want)
	})

	t.Run("English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, i0Ek3"

		assertMsg(t, got, want)
	})

}
