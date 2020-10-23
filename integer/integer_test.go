package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 1)
	want := 3

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
