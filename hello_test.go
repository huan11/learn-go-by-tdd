package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Asher"
	got := Hello("Asher")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
