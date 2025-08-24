package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "hahaha"
	got := Repeat("ha")
	if got != want {
		t.Errorf("want %q; got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("ha")
	}
}
