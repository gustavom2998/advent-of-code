package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input2.txt
var input1 string

func main() {
	segments := strings.Split(input1, "\n")
	fmt.Println(FindMultiples(segments))
}

func FindMultiples(segments []string) int {
	totalSum := 0
	for _, s := range segments {
		m, err := Multiples(s)
		if err != nil {
			fmt.Printf("unable to process segment: %v\n", err.Error())
			return 0
		}

		sum := 0
		for _, val := range m {
			sum += val[0] * val[1]
		}
		totalSum += sum
	}
	return totalSum
}

type State uint8

const (
	Start State = iota
	MulWithPar
	FirstOperand
	SecondOperand
	Comma
	ClosePar
)

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
