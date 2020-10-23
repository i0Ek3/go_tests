package iteration

import (
	"fmt"
	"testing"
)

const (
	count = 3
	str   = "a"
)

func TestIterate(t *testing.T) {
	got := Repeat(count, str)
	want := "aaa"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(count, str)
	}
}

func ExampleRepeat() {
	ret := Repeat(count, str)
	fmt.Println(ret)
	// Output: aaa
}
