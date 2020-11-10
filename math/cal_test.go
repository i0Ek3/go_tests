package main

import (
	"math"
	"testing"

	a "github.com/i0Ek3/asrt"
)

func TestCal(t *testing.T) {
	t.Run("-0.5", func(t *testing.T) {
		number := -0.5
		got := Cal(number)
		want := math.Cos(number)

		a.Asrt(t, got, want)
	})

	t.Run("0.5", func(t *testing.T) {
		number := 0.5
		got := Cal(number)
		want := math.Cos(number)

		a.Asrt(t, got, want)
	})

	t.Run("0", func(t *testing.T) {
		number := 0.0
		got := int(Cal(number))
		want := 1

		a.Asrt(t, got, want)
	})
}

func BenchmarkCal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cal(math.Pi)
	}
}
