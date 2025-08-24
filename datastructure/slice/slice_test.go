package slice

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("Sum multiple slice ", func(t *testing.T) {
		want := []int{5, 5}
		got := SumAll([]int{2, 3}, []int{2, 3})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}

	})
}
func TestSumAllTails(t *testing.T) {
	want := []int{4, 6}
	got := SumAllTails([]int{1, 2, 2}, []int{2, 3, 3})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}
