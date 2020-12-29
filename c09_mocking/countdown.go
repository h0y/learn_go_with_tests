package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func Countdown(out io.Writer, s Sleeper) {
	for i:= 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	s.Sleep()
	fmt.Fprint(out, "go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}