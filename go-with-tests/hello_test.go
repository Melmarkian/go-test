package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Matthias")
	want := "Hello, Matthias"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
