package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	main "github.com/gustavom2998/advent-of-code/2024/03"
)

func TestExerciseA(t *testing.T) {
	type MemorySegmentCase struct {
		Input                   []string
		WantTotal, WantFiltered int
	}
	segments := []MemorySegmentCase{
		MemorySegmentCase{
			Input:        []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			WantTotal:    161,
			WantFiltered: 8,
		},
        MemorySegmentCase{
			Input:        []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			WantTotal:    161,
			WantFiltered: 48,
		},

	}

	for i, segment := range segments {
		gotTotal, gotFiltered := main.FindMultiples(segment.Input)
		wantTotal, wantFiltered := segment.WantTotal, segment.WantFiltered
		if diff := cmp.Diff(wantTotal, gotTotal); diff != "" {
			t.Errorf("Segment %v - findMultiples() total mismatch (-want +got):\n%s", i, diff)
		}
		if diff := cmp.Diff(wantFiltered, gotFiltered); diff != "" {
			t.Errorf("Segment %v - findMultiples() filtered mismatch (-want +got):\n%s", i, diff)
		}
	}

}
