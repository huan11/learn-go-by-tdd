package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Asher"
		got := Hello("Asher")
		assertCorrectMessage(t, want, got)
	})

	t.Run("say 'Hello, world' when an empty string is supplied ", func(t *testing.T) {
		want := "Hello, world2"
		got := Hello("")
		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t *testing.T, want string, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
