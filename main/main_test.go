package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{write, sleep, write, sleep, write, sleep, write}
		got := spySleepPrinter.Calls

		if !reflect.DeepEqual(got, want) { //why did i have to use deepEqual here?
			t.Errorf("Got calls %v wanted %v", got, want)
		}
	})

}

//with this, we have the function and its two most important properties tested
// i.e. we have the writing tested, and the order of sleep and write testied

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

//essentially, the configurable sleeper is not a sleeper function in itself. It is a struct with a time field and a method which sleeps for the time duration
// we have a spy function that provides the sleeping functionality.
// so we can create a configurable sleeper with a time and a spy function.
// calling the spy function with the time should record the time it was called with to the s.durationslept field so you can see the number it was called with
// a brilliant way to spy on something!
// this must mean that we can use the configurable sleeper in main if we just pass in time.Sleep? - to check
