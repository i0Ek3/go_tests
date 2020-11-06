package rnr

import (
	"strings"
)

// RNR defines a map: Value to Roman
type RNR struct {
	Value uint16
	Roman string
}

type RomanNumber []RNR
type windowedRoman string

// RN returns a corresponding string with given number
func RN(number uint16) string {
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
func NR(roman string) (ret uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		ret += mapRNR.ValueOf(symbols...)
	}
	return
}

// ValueOf fetchs value from given arguments
func (r RomanNumber) ValueOf(romans ...byte) uint16 {
	roman := string(romans)
	for _, s := range r {
		if s.Roman == roman {
			return s.Value
		}
	}
	return 0
}

// Exists checks if roman equal Roman
func (r RomanNumber) Exists(romans ...byte) bool {
	roman := string(romans)
	for _, s := range r {
		if s.Roman == roman {
			return true
		}
	}
	return false
}

// mapRNR defines a map for RNR
var mapRNR = RomanNumber{
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

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		if i+1 < len(w) && isSubtractive(symbol) && mapRNR.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
