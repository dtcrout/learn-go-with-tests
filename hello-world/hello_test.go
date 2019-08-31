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

	// Assertion message for subtasks wrapped into a function
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite the function is a helper
		// If this is commented out, the line number reported back will be this
		// if statement instead of where it actually fails
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Subtest for argument in Hello()
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Darshan")
		want := "Hello, Darshan"

		assertCorrectMessage(t, got, want)
	})

	// Subtest for empty argument in Hello()
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
