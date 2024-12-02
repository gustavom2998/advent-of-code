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
	a, b := ProcessPairs(input1)
	if len(a) > 0 && len(b) > 0 {
		// Answer: 765748, 27732508
		fmt.Printf("Input 1 - Distance: %v Similarity: %v\n", Distance(a, b), Similarity(a, b))
	}
	c, d := ProcessPairs(input2)
	if len(c) > 0 && len(d) > 0 {
		// Answer: 765748, 27732508
		fmt.Printf("Input 2 - Distance: %v Similarity: %v\n", Distance(c, d), Similarity(c, d))
	}
}

// ProcessPairs receives a string file and returns array pairs.
//
// It parses the file, breaking it down by columns (col a and col b)
// and returns N items in each array, where each item is an int.
func ProcessPairs(pairFile string) ([]int, []int) {
	var err error
	var val int
	a := []int{}
	b := []int{}

	for _, row := range strings.Split(strings.TrimSpace(pairFile), "\n") {
		rowSlice := strings.Fields(row)
		if len(rowSlice) == 0 {
			return nil, nil
		}
		val, err = strconv.Atoi(rowSlice[0])
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
			return nil, nil
		}
		a = append(a, val)
		val, err = strconv.Atoi(rowSlice[1])
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
			return nil, nil
		}
		b = append(b, val)
	}

	return a, b
}

// Sort orders an integer slice using insertion sort
func Sort(a []int) []int {
	for i := range a {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}

	return a
}

// Distance orders the two slices and returns the difference between each element
func Distance(a, b []int) (result int) {
	a = Sort(a)
	b = Sort(b)

	for i, v1 := range a {
		v2 := b[i]
		value := v1 - v2
		if value < 0 {
			value = -1 * value
		}
		result += value
	}

	return result
}

// Similarity sums the number on the left list by the amount of times they happen on the right list
func Similarity(a, b []int) (result int) {
	count := map[int]int{}
	for _, v1 := range a {
		for _, v2 := range b {
			if v2 == v1 {
				count[v1] += 1
			}
		}
	}

	for k, v := range count {
		result += k * v
	}
	return result
}
