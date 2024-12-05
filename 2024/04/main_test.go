package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	main "github.com/gustavom2998/advent-of-code/2024/04"
)

func TestExerciseA(t *testing.T) {
	type CrosswordCase struct {
		Input     string
		WantTotal int
		WantCross int
	}
	cases := []CrosswordCase{
		CrosswordCase{
			Input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			WantTotal: 18,
			WantCross: 9,
		},
	}

	for i, c := range cases {
		gotTotal, gotCross := main.CountXMAS(c.Input)
		wantTotal := c.WantTotal
		wantCross := c.WantCross
		if diff := cmp.Diff(wantTotal, gotTotal); diff != "" {
			t.Errorf("Crossword %v - CountXMAS() total mismatch (-want +got):\n%s", i, diff)
		}
		if diff := cmp.Diff(wantCross, gotCross); diff != "" {
			t.Errorf("Crossword %v - CountXMAS() cross mismatch (-want +got):\n%s", i, diff)
		}
	}

}
