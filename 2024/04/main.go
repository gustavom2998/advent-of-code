package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input2.txt
var input string

func main() {
	fmt.Println(CountXMAS(input))
}

// CountXMAS can be used to do a word search for xmas and return the count
func CountXMAS(s string) (totalSum int, crossSum int) {
	matrix := [][]string{}
	for _, row := range strings.Split(strings.TrimSpace(s), "\n") {
		matrix = append(matrix, strings.Split(strings.ToLower(row), ""))
	}

	for i, row := range matrix {
		for j, col := range row {
			if col == "x" {
				// Check left to right
				if (j <= (len(row) - 4)) && matrix[i][j+1] == "m" && matrix[i][j+2] == "a" && matrix[i][j+3] == "s" {
					totalSum++
				}
				// Check right to left
				if ((j - 3) >= 0) && matrix[i][j-1] == "m" && matrix[i][j-2] == "a" && matrix[i][j-3] == "s" {
					totalSum++
				}
				if (i - 3) >= 0 {
					// Check up
					if matrix[i-1][j] == "m" && matrix[i-2][j] == "a" && matrix[i-3][j] == "s" {
						totalSum++
					}
					// Check up left to right
					if (j <= (len(row) - 4)) && matrix[i-1][j+1] == "m" && matrix[i-2][j+2] == "a" && matrix[i-3][j+3] == "s" {
						totalSum++
					}
					// Check up right to left
					if ((j - 3) >= 0) && matrix[i-1][j-1] == "m" && matrix[i-2][j-2] == "a" && matrix[i-3][j-3] == "s" {

						totalSum++
					}
				}
				if i <= (len(matrix) - 4) {
					// Check down
					if matrix[i+1][j] == "m" && matrix[i+2][j] == "a" && matrix[i+3][j] == "s" {
						totalSum++
					}
					// Check down left to right
					if (j <= (len(row) - 4)) && matrix[i+1][j+1] == "m" && matrix[i+2][j+2] == "a" && matrix[i+3][j+3] == "s" {
						totalSum++
					}
					// Check down right to left
					if ((j - 3) >= 0) && matrix[i+1][j-1] == "m" && matrix[i+2][j-2] == "a" && matrix[i+3][j-3] == "s" {
						totalSum++
					}
				}
			} else if matrix[i][j] == "a" {
				if i < (len(matrix)-1) && ((i - 1) >= 0) && (j < (len(row) - 1)) && ((j - 1) >= 0) {
					if matrix[i-1][j-1] == "m" && matrix[i+1][j+1] == "s" {
						if matrix[i-1][j+1] == "m" && matrix[i+1][j-1] == "s" {
							crossSum++
						}
						if matrix[i-1][j+1] == "s" && matrix[i+1][j-1] == "m" {
							crossSum++
						}
					}
					if matrix[i-1][j-1] == "s" && matrix[i+1][j+1] == "m" {
						if matrix[i-1][j+1] == "m" && matrix[i+1][j-1] == "s" {
							crossSum++
						}
						if matrix[i-1][j+1] == "s" && matrix[i+1][j-1] == "m" {
							crossSum++
						}
					}
				}
			}

		}
	}
	return
}
