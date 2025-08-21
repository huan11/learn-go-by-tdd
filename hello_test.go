package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Asher"
	got := Hello("Asher")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Asher"
		got := Hello("Asher")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, world' when an empty string is supplied ", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
