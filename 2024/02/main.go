package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input1.txt
var input1 string

//go:embed input2.txt
var input2 string

func main() {
	a := ProcessReports(input1)
	if len(a) > 0 {
		// Answer: 2
		fmt.Printf("Input 1 - : %v\n", CountSafe(a))
	}
	b := ProcessReports(input2)
	if len(b) > 0 {
		// Answer: 2
		fmt.Printf("Input 2 - : %v\n", CountSafe(b))
	}
}

// ProcessReports receives a string file and returns an array of reports
//
// It parses the file, breaking it down a report (int array) per row
func ProcessReports(reportFile string) [][]int {
	a := [][]int{}

	for _, row := range strings.Split(strings.TrimSpace(reportFile), "\n") {
		reportStr := strings.Fields(row)
		if len(reportStr) == 0 {
			return nil
		}
		report := []int{}
		for _, val := range reportStr {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Failed to process report: %v", err.Error())
				return nil
			}
			report = append(report, intVal)
		}
		a = append(a, report)
	}

	return a
}

// CountSafe receives a slice of reports and counts the number of safe reports
func CountSafe(a [][]int) (count int) {
	for _, r := range a {
		if CheckSequence(r) {
			count += 1
		}
	}
	return
}

// CheckSequence checks if a sequence is safe.
//
// A sequence is safe if it constantly increases or decreases.
// Each value cannot increase by less than 1 and more than 3 compared to the previous.
func CheckSequence(a []int) (result bool) {
	if len(a) <= 1 {
		return true
	}
	result = true
	lastDiff := 1
	for i := 1; i < len(a); i++ {
		currentDiff := a[i] - a[i-1]
		// Check constantly increasing or decreasing
		if (i > 1) && ((lastDiff > 0) != (currentDiff > 0)) {
			return false
		}

		// Check increase or decrease is greater than zero and less than four
		if abs(currentDiff) < 1 || abs(currentDiff) > 3 {
			return false
		}

		lastDiff = currentDiff
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
