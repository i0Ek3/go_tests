package smi

import (
	"math"
)

// Shape defines some shapes
type Shape struct {
	length float64
	width  float64
}

// Shaper defines some method of Shape
type Shaper interface {
	Perimeter() float64
	Area() float64
}

// Rectangle
type Rectangle struct {
	Shape
}

// Circle
type Circle struct {
	radius float64
}

// Triangle
type Triangle struct {
	length float64
	height float64
}

// Perimeter calculates the perimeter of given shape
//func (r Rectangle) Perimeter() float64 {
//	r.area = r.length * r.width
//	return r.area
//}

//// Perimeter calculates the perimeter of given shape
//func Perimeter(l, w float64) float64 {
//	return 2 * (l + w)
//}
//
//// Area calculates the area of given shape
//func Area(l, w float64) float64 {
//	return l * w
//}

// Perimeter calculates the perimeter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Area calculates the area of Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Perimeter calculates the perimeter of Circle
func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// Area calculates the area of Circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(2, c.radius)
}

// Perimeter calculates the area of Triangle
func (t Triangle) Perimeter() float64 {
	// cannot calculate the perimeter of Triangle directly
	return 0.0
}

// Area calculates the area of Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.length * t.height
}
