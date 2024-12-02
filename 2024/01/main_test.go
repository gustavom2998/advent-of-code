package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	main "github.com/gustavom2998/advent-of-code/2024/01"
)

func TestSort(t *testing.T) {
	type testCase struct {
		Input []int
		Want  []int
	}

	testCases := []testCase{
		testCase{
			Input: []int{3, 6, 5, 8, 5, 9, 2, 1},
			Want:  []int{1, 2, 3, 5, 5, 6, 8, 9},
		},
		testCase{
			Input: []int{3, 1, 2, 4},
			Want:  []int{1, 2, 3, 4},
		},
		testCase{
			Input: []int{4, 3, 5, 2, 1},
			Want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		want := tc.Want
		got := main.Sort(tc.Input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Sort() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestExercise(t *testing.T) {
	a := main.Sort([]int{3, 4, 2, 1, 3, 3})
	b := main.Sort([]int{4, 3, 5, 3, 9, 3})
	got := main.Distance(a, b)
	want := 11
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Distance() mismatch (-want +got):\n%s", diff)
	}
	got = main.Similarity(a, b)
	want = 31
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Similarity() mismatch (-want +got):\n%s", diff)
	}
}
