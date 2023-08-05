package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Gabriel")
  want := "Hello, Gabriel"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
