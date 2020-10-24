package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	check := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got '%q' want '%d'", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		check(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		check(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	check := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		num1 := []int{1, 2, 3, 4, 5, 6}
		num2 := []int{0, 8, 3}

		got := SumAll(num1, num2)
		want := 32

		check(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		num1 := []int{6}
		num2 := []int{0, 3}

		got := SumAll(num1, num2)
		want := 9

		check(t, got, want)
	})
}

func TestSumToNew(t *testing.T) {
	check := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		got := SumToNew([]int{1, 2, 3}, []int{0, 1})
		want := []int{6, 1}
		check(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := SumToNew([]int{3}, []int{})
		want := []int{3, 0}
		check(t, got, want)
	})
}

func ExampleSum() {
	nums := []int{1, 2, 3, 4}
	ret := Sum(nums)
	fmt.Println(ret)
	// Output: 10
}

func ExampleSumAll() {
	num1 := []int{2, 3}
	num2 := []int{1, 2, 3, 4}
	ret := SumAll(num1, num2)
	fmt.Println(ret)
	// Output: 15
}

func ExampleSumToNew() {
	num1 := []int{}
	num2 := []int{3, 3}
	newslice := SumToNew(num1, num2)
	fmt.Println(newslice)
	// Output: [0 6]
}
