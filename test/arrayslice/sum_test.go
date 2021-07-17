package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumForArray(numbers)
	want := 15

	if got != want {
		t.Errorf("get %d want %d given, %v", got, want, numbers)
	}
}

func TestSum2(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumForSlice(numbers)
		want := 15

		if got != want {
			t.Errorf("get %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumForSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("get %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// if got != want {
	// 	t.Errorf("got %v wangt %v", got, want)
	// }

	// reflect.DeepEqual is a convenient way of comparing slices (and other things)
	// you must be careful when using it.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wangt %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails2(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSumAllTails3(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		// Helper marks the calling function as a test helper function.
		// When printing file and line information, that function will be skipped.
		// Helper may be called simultaneously from multiple goroutines.
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