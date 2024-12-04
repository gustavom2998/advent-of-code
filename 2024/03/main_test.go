package main_test

import (
	"testing"

	main "github.com/gustavom2998/advent-of-code/2024/03"
	"github.com/google/go-cmp/cmp"
)

func TestExerciseA(t *testing.T) {
	type MemorySegmentCase struct {
		Input []string
		Want  int
	}
	segments := []MemorySegmentCase{
		MemorySegmentCase{
			Input: []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			Want:  161,
		},
	}

	for i, segment := range segments {
		got := main.FindMultiples(segment.Input)
		want := segment.Want
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Segment %v - findMultiples() mismatch (-want +got):\n%s", i, diff)
		}
	}

}
