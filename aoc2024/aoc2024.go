package aoc2024

import (
	"fmt"
	"github.com/fundthmcalculus/advent-of-code/helpers"
	"strconv"
	"strings"
)

func Problem7() {
	lines := helpers.LoadInputLines(2024, 7, false)
	totalOutcomes := int64(0)
	totalOutcomes2 := int64(0)
	for _, line := range lines {
		// Split at the spaces, drop the ":"
		numbers := helpers.ToInt64(strings.Split(strings.Replace(line, ":", "", 1), " ")...)
		outcome := evaluate2(numbers, false)
		outcome2 := evaluate2(numbers, true)
		totalOutcomes += outcome
		totalOutcomes2 += outcome2
	}

	fmt.Println("Problem 7A:", totalOutcomes)
	fmt.Println("Problem 7B:", totalOutcomes2)
}

func evaluate2(numbers []int64, concatenate bool) int64 {
	target := numbers[0]
	// Generate all possible permutations of operators
	operators := make([]byte, len(numbers)-2)
	// Fill with "+", then keep changing them
	for i := 0; i < len(operators); i++ {
		operators[i] = '+'
	}
	for true {
		// Evaluate this option
		res := eval(numbers, operators)
		if res == target {
			return res
		}
		// index the first operator, then carry as needed
		carry := true
		for i := 0; i < len(operators); i++ {
			if !carry {
				break
			}
			carry = false
			if operators[i] == '+' {
				if !concatenate {
					operators[i] = '*'
				} else {
					operators[i] = '&'
				}
			} else if operators[i] == '*' {
				carry = true
				operators[i] = '+'
			} else if operators[i] == '&' {
				operators[i] = '*'
			}
		}
		// If we carried over the end, we're done.
		if carry {
			return 0
		}
	}
	return 0
}

func eval(numbers []int64, operators []byte) int64 {
	runningTotal := numbers[1]
	for i, operator := range operators {
		if runningTotal > numbers[0] {
			return 0
		}
		x2 := numbers[i+2]
		if operator == '+' {
			runningTotal += x2
		} else if operator == '*' {
			runningTotal *= x2
		} else if operator == '&' {
			// Concatenate strings because lazy
			p1 := strconv.Itoa(int(runningTotal))
			p2 := strconv.Itoa(int(x2))
			runningTotal = helpers.ToInt64(p1 + p2)[0]
		}
	}
	return runningTotal
}
