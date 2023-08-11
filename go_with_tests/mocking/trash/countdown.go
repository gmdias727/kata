package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

type DefaultSleeper struct{}

type Sleeper interface {
	Sleep()
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
