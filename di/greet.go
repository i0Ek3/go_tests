package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet greets someone
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s", name)
}

// GreetHandler handles request and ack response
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "i0Ek3")
}

func main() {
	fmt.Println("Please open the url http://localhost:6789 to visit it.")
	http.ListenAndServe(":6789", http.HandlerFunc(GreetHandler))
}
