package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Blake")
  want := "Hello, Blake"

  if (got != want) {
    t.Errorf("got %q want %q", got, want)
  }
}
