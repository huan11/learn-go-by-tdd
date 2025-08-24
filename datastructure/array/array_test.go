package array

import "testing"

func TestSum(t *testing.T) {
	numbers := [3]int{10, 20, 20}

	want := 50
	got := Sum(numbers)

	if want != got {
		t.Errorf("want %d, got %d, %v", want, got, numbers)
	}
}
