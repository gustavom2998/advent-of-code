package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input2.txt
var input string

func main() {
	fmt.Println(FindMultiples(input))
}

// FindMultiples can be used to get the multiples total sum and filtered sum
func FindMultiples(s string) (totalSum int, filteredSum int) {
	m, err := Multiples(s)
	if err != nil {
		fmt.Printf("unable to process segment: %v\n", err.Error())
		return 0,0
	}

	sum := 0
	for _, val := range m {
		sum += val[0] * val[1]
	}
	totalSum += sum

	m, err = FilteredMultiples(s)
    fmt.Println(m)
	if err != nil {
		fmt.Printf("unable to process segment: %v\n", err.Error())
		return 0,0
	}

	sum = 0
	for _, val := range m {
		sum += val[0] * val[1]
	}
	filteredSum += sum
	
	return
}

// State is used to represent the states used to identify multiples
type State uint8

const (
	Start State = iota  // Start is the initial state
	MulWithPar          // MulWithPar is the State when we've matched a mul command and need to find the operands
	FirstOperand // FirstOperand is the state used to identify the digits for the first operand
	SecondOperand // SecondOperand is the state used after identifying a comma after the first operand
)
// Filtered multiples returns the multiple pairs check for do(). 
//
// This could be simplified to use a regex but I wanted to have fun with a Finite State Machine.
func FilteredMultiples(str string) ([][]int, error) {
	state := Start
	pairs := [][]int{}
	s := strings.Split(str, "")
	numbers1 := []string{}
	numbers2 := []string{}

	do := true
	for i := 0; i < len(s); {
		switch state {
		case Start:
			if i < (len(s)-4) && slices.Equal(s[i:i+4], []string{"m", "u", "l", "("}) && do {
				state = MulWithPar
				i += 4
			} else if i < (len(s)-4) && slices.Equal(s[i:i+4], []string{"d", "o", "(", ")"}) {
				do = true
				i += 4
			} else if i < (len(s)-7) && slices.Equal(s[i:i+7], []string{"d", "o", "n", "'", "t", "(", ")"}) {
                do = false
				i += 7
			} else {
				i++
			}
		case MulWithPar:
			if isDigit(s[i]) {
				state = FirstOperand
			}
		case FirstOperand:
			if isDigit(s[i]) {
				numbers1 = append(numbers1, s[i])
				i++
			} else if s[i] == "," && len(numbers1) > 0 {
				state = SecondOperand
				i++
			} else {
				numbers1 = []string{}
				state = Start
			}
		case SecondOperand:
			if isDigit(s[i]) {
				numbers2 = append(numbers2, s[i])
				i++
			} else if s[i] == ")" && len(numbers2) > 0 {
				first, err := strconv.Atoi(strings.Join(numbers1, ""))
				if err != nil {
					return nil, err
				}
				second, err := strconv.Atoi(strings.Join(numbers2, ""))
				if err != nil {
					return nil, err
				}
				pairs = append(pairs, []int{first, second})
				numbers1 = []string{}
				numbers2 = []string{}
				i++
				state = Start
			} else {
				numbers1 = []string{}
				numbers2 = []string{}
				state = Start
			}
		default:
			return nil, fmt.Errorf("invalid state")
		}
	}
	return pairs, nil
}

// Filtered multiples returns the multiple pairs. 
//
// This could be simplified to use a regex but I wanted to have fun with a Finite State Machine.
func Multiples(str string) ([][]int, error) {
	state := Start
	pairs := [][]int{}
	s := strings.Split(str, "")
	numbers1 := []string{}
	numbers2 := []string{}
	for i := 0; i < len(s); {
		switch state {
		case Start:
			if i < (len(s)-4) && slices.Equal(s[i:i+4], []string{"m", "u", "l", "("}) {
				state = MulWithPar
				i += 4
			} else {
				i++
			}
		case MulWithPar:
			if isDigit(s[i]) {
				state = FirstOperand
			}
		case FirstOperand:
			if isDigit(s[i]) {
				numbers1 = append(numbers1, s[i])
				i++
			} else if s[i] == "," && len(numbers1) > 0 {
				state = SecondOperand
				i++
			} else {
				numbers1 = []string{}
				state = Start
			}
		case SecondOperand:
			if isDigit(s[i]) {
				numbers2 = append(numbers2, s[i])
				i++
			} else if s[i] == ")" && len(numbers2) > 0 {
				first, err := strconv.Atoi(strings.Join(numbers1, ""))
				if err != nil {
					return nil, err
				}
				second, err := strconv.Atoi(strings.Join(numbers2, ""))
				if err != nil {
					return nil, err
				}
				pairs = append(pairs, []int{first, second})
				numbers1 = []string{}
				numbers2 = []string{}
				i++
				state = Start
			} else {
				numbers1 = []string{}
				numbers2 = []string{}
				state = Start
			}
		default:
			return nil, fmt.Errorf("invalid state")
		}
	}
	return pairs, nil
}

func isDigit(c string) bool {
	switch c {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}
