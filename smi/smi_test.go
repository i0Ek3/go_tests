package smi

import (
	"fmt"
	//"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertTest := []struct {
		name  string
		shape Shaper
		want  float64
	}{
		//{"R", Rectangle{Shape{1.0, 20.0}}, 42.0},
		//{"C", Circle{2.0}, 12.566370614359172},
		{name: "Rectangle", shape: Rectangle{Shape{length: 1.0, width: 20.0}}, want: 42.0},
		{name: "Circle", shape: Circle{radius: 2.0}, want: 12.566370614359172},
	}

	for _, tt := range assertTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("got '%f' want '%f'", got, tt.want)
			}
		})
	}
}

/*
func TestPerimeter(t *testing.T) {
	check := func(t *testing.T, shape Shaper, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Shape{1.0, 20.0}}
		//got := r.Perimeter()
		want := 42.0
		check(t, r, want)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{2.0}
		//got := c.Perimeter()
		want := 12.566370614359172
		check(t, c, want)
	})
}
*/

func ExamplePerimeter() {
	//r := Rectangle{Shape{1.0, 20.0}}
	//ret := r.Perimeter()
	//fmt.Println(ret)
	//// Output: 42

	c := Circle{2.0}
	ret := c.Perimeter()
	fmt.Println(ret)
	// Output: 12.566370614359172
}

func TestArea(t *testing.T) {
	assertTest := []struct {
		name  string
		shape Shaper
		want  float64
	}{
		//{"R", Rectangle{Shape{1.0, 20.0}}, 20.0},
		//{"C", Circle{2.0}, 12.566370614359172},
		//{"T", Triangle{2.0, 1.0}, 1.0},
		{name: "Rectangle", shape: Rectangle{Shape{length: 1.0, width: 20.0}}, want: 20.0},
		{name: "Circle", shape: Circle{radius: 2.0}, want: 12.566370614359172},
		{name: "Triangle", shape: Triangle{length: 2.0, height: 1.0}, want: 1.0},
	}

	for _, tt := range assertTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got '%f' want '%f'", got, tt.want)
			}
		})
	}
}

/*
func TestArea(t *testing.T) {
	check := func(t *testing.T, shape Shaper, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Shape{1.0, 20.0}}
		//got := r.Area()
		want := 20.0
		check(t, r, want)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{2.0}
		//got := c.Area()
		want := 4.0 * math.Pi
		check(t, c, want)
	})
}
*/

func ExampleArea() {
	//r := Rectangle{Shape{1.0, 20.0}}
	//ret := r.Area()
	//fmt.Println(ret)
	//// Output: 20

	c := Circle{2.0}
	ret := c.Area()
	fmt.Println(ret)
	// Output: 12.566370614359172
}
