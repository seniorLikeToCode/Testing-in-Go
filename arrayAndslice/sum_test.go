package array_slice

import (
	"reflect"
	"testing"
)

func assertSum(t testing.TB, got int, hasSum int, nums []int) {
	t.Helper()
	if got != hasSum {
		t.Errorf("got %d want %d given, %v", got, hasSum, nums)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		hasSum := 15

		assertSum(t, got, hasSum, nums)
	})

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		hasSum := 6

		assertSum(t, got, hasSum, nums)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of array sum", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}