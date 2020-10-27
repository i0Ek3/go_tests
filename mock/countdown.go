package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	begin = 3

	last  = "Go!"
	sleep = "sleep"
	write = "write"
)

// Sleeper defines sleep interface
type Sleeper interface {
	Sleep()
}

// Spy is a mock to record how many Sleep() was invocated
type SpySleep struct {
	Calls int
}

// Sleep mocks sleep
func (s *SpySleep) Sleep() {
	s.Calls++
}

// CountdownOperationSpy spys times of call
type CountdownOperationSpy struct {
	Calls []string
}

// Sleep add sleep into slice
func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write writes into slice
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ConfigurableSleeper defines sleep duration
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep sleeps duration time
func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.duration)
}

// Countdown counts the number
func Countdown(out io.Writer, slp Sleeper) {
	for i := begin; i > 0; i-- {
		slp.Sleep()
		fmt.Fprintln(out, i)
	}
	slp.Sleep()
	fmt.Fprint(out, last)
}

func main() {
	slp := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, slp)
}
