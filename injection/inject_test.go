package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nkechi")

	got := buffer.String()
	want := "Hello, Nkechi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
