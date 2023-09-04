package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

type Sleeper interface { //creating an interface to enable us mock our time.Sleep function
	Sleep()
}

type SpySleeper struct { //creating a spy to record how many times our sleep function is called without actually calling it in the test
	Calls int
}

func (s *SpySleeper) Sleep() { // a function to satisfy the Sleeper interface
	s.Calls++
}

type SpyCountdownOperations struct { // creating a combined spy to check that the order of operations is right
	Calls []string
}

// the below satisfies the Sleeper interface
func (s *SpyCountdownOperations) Sleep() { // I think you use a pointer on a receiver function when you want that particular item being modified? To check
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) { //satisfies the writer interface as well
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

//Essentially, above should show the correct order of sequences, which is write, sleep, write, sleep, write, sleep, write. This should help us test that our order is right

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) //includes the function signature of time.Sleep
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
