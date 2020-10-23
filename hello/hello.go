package main

import (
	"fmt"
)

const (
	greetingEN = "Hello, "
	greetingSP = "Hola, "
)

// Hello says hello for you
func Hello(s, lan string) string {
	if s == "" {
		s = "i0Ek3"
	}
	return greeting(lan) + s
}

func greeting(lan string) (prefix string) {
	switch lan {
	case "Espanol":
		prefix = greetingSP
	case "English":
		prefix = greetingEN
	default:
	}

	return
}

func main() {
	fmt.Println(Hello("i0Ek3", "Espanol"))
}
