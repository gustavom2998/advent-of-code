package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	main "github.com/gustavom2998/advent-of-code/2024/02"
)

func TestExerciseA(t *testing.T) {
	type ReportCase struct {
		Input []int
		Want  bool
	}
	reports := []ReportCase{
		ReportCase{
			Input: []int{7, 6, 4, 2, 1},
			Want:  true,
		},
		ReportCase{
			Input: []int{1, 2, 7, 8, 9},
			Want:  false,
		},
		ReportCase{
			Input: []int{9, 7, 6, 2, 1},
			Want:  false,
		},
		ReportCase{
			Input: []int{1, 3, 2, 4, 5},
			Want:  false,
		},
		ReportCase{
			Input: []int{8, 6, 4, 4, 1},
			Want:  false,
		},
		ReportCase{
			Input: []int{1, 3, 6, 7, 9},
			Want:  true,
		},
	}

	for i, report := range reports {
		got, _ := main.CheckSequence(report.Input)
		want := report.Want
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Report %v - CheckReport() mismatch (-want +got):\n%s", i, diff)
		}
	}
}

func TestExerciseB(t *testing.T) {
	type ReportCase struct {
		Input []int
		Want  bool
	}
	reports := []ReportCase{
		ReportCase{
			Input: []int{7, 6, 4, 2, 1},
			Want:  true,
		},
		ReportCase{
			Input: []int{1, 2, 7, 8, 9},
			Want:  false,
		},
		ReportCase{
			Input: []int{9, 7, 6, 2, 1},
			Want:  false,
		},
		ReportCase{
			Input: []int{1, 3, 2, 4, 5},
			Want:  true,
		},
		ReportCase{
			Input: []int{8, 6, 4, 4, 1},
			Want:  true,
		},
		ReportCase{
			Input: []int{1, 3, 6, 7, 9},
			Want:  true,
		},
		ReportCase{
			Input: []int{6, 7, 9, 11, 13, 14, 18},
			Want:  true,
		},
		ReportCase{
			Input: []int{1, 5, 6, 7},
			Want:  true,
		},
		ReportCase{
			Input: []int{81, 82, 81, 80, 79},
			Want:  true,
		},
	}

	for i, report := range reports {
		got := main.CheckSequenceWithFailure(report.Input)
		want := report.Want
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Report %v - CheckReport() mismatch (-want +got):\n%s", i, diff)
		}
	}
}
