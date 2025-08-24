package slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum the collection of any numbers", func(t *testing.T) {
		numbers := []int{2, 2, 2, 2, 2}

		want := 10
		got := Sum(numbers)

		if got != want {
			t.Errorf("want %d, got %d, %v", want, got, numbers)
		}
	})
}
