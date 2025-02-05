package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sum of slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{6, 5, 4, 3, 2, 1})
		want := []int{2, 9, 15}

		checkSums(t, got, want)
	})

	t.Run("safely handly nil slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{21, 12, 4})
		want := []int{0, 16}

		checkSums(t, got, want)
	})
}
