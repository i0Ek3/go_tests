package array

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	num := [5]int{1, 2, 3, 4, 5}

	got := Sum2(num)
	want := 15

	if got != want {
		t.Errorf("got '%q' want '%d' given arr '%v'", got, want, num)
	}
}

func ExampleSum() {
	num := [...]int{1, 2, 3, 4, 5}
	ret := Sum2(num)
	fmt.Println(ret)
	// Output: 15
}
