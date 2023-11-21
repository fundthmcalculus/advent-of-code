package aoc2015

import (
	"fmt"
	"os"
	"path"
)

const aoc2015 = "aoc2015"

func loadInput(day int) string {
	b, err := os.ReadFile(path.Join(aoc2015, string(rune(day))+".txt"))
	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

public func Problem1() {
	// Load p1.txt
	lispData := loadInput(1)
	curFloor := 0
	for i := 0; i < len(lispData); i++ {
		if lispData[i] == '(' {
			curFloor++
		} else if lispData[i] == ')' {
			curFloor--
		}
	}
	fmt.Println("Problem 1 Answer:", curFloor)
}
