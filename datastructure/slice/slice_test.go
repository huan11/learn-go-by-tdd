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

	t.Run("Should panic when using fixed-size slice for Sum2", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic due to index out of bounds")
			}
		}()

		// SumAll2 使用固定大小为2的slice，但传入3个参数会导致越界
		SumAll2([]int{1, 2, 2}, []int{2, 3, 3}, []int{2})
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	}

	t.Run("Sum multiple slice ", func(t *testing.T) {
		want := []int{4, 6}
		got := SumAllTails([]int{1, 2, 2}, []int{2, 3, 3})

		checkSums(t, got, want)
	})

	t.Run("handle empty slice", func(t *testing.T) {
		want := []int{0, 6}
		got := SumAllTails([]int{}, []int{2, 3, 3})

		checkSums(t, got, want)
	})

}
