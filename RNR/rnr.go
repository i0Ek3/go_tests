package main

import (
	"strings"
)

// RNR defines a map: Value to Roman
type RNR struct {
	Value int
	Roman string
}

// mapRNR defines a map for RNR
var mapRNR = []RNR{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// RN returns a corresponding string with given number
func RN(number int) string {
	var ret strings.Builder
	for _, v := range mapRNR {
		for number >= v.Value {
			ret.WriteString(v.Roman)
			number -= v.Value
		}
	}
	return ret.String()
}

// NR returns a corresponding int with given roman
func NR(roman string) int {
	ret := 0
	for range roman {
		ret++
	}
	return ret
}
