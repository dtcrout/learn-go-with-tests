/*
Test function for Hello().

Conventions:
  - Test files start have names like xxx_test.go
  - Test functions must start with word Testing
  - The test function only takes one argument only: t *testing.T
*/

package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello()
	got := Hello("Darshan")
	// want := "Hello, world"
	// want := "Hello world"
	want := "Hello, Darshan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
