package main

import (
	"math"
)

// Cal calculates the value for given needs
func Cal(number float64) float64 {
	if number < 0 {
		return mathHelper(number)
	} else if number == 0 {
		math.Inf(0)
	}
	return math.Cos(number)
}

func mathHelper(number float64) float64 {
	return calNegative(number)
}

func calNegative(number float64) float64 {
	return math.Cos(math.Abs(number))
}
