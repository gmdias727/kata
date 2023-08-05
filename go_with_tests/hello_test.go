package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel", "")
		want := "Hello, Gabriel"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Gabriel", "French")
		want := "Bonjour, Gabriel"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// This function is a helper to avoid repeating the same code in every test.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
