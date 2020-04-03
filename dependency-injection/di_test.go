package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// Buffer location and string
	Greet(&buffer, "Darshan")

	// String returns the contents of the unread portion of the buffer
	// as a string
	got := buffer.String()
	want := "Hello, Darshan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
