package aoc2023

import (
	"fmt"
	"github.com/fundthmcalculus/advent-of-code/helpers"
	"strconv"
	"strings"
	"unicode"
)

func makeEmptyBag() map[string]int64 {
	bag := make(map[string]int64)
	bag["red"] = 0
	bag["green"] = 0
	bag["blue"] = 0
	return bag
}

func Problem2() {
	bagLoad := make(map[string]int64)
	bagLoad["red"] = 12
	bagLoad["green"] = 13
	bagLoad["blue"] = 14
	lines := helpers.LoadInputLines(2023, 2, false)
	impossibleGameIdSum := 0
	possibleGameIdSum := 0
	gamePowerSum := 0
	for _, line := range lines {
		minGameBag := makeEmptyBag()
		gamePossible := true
		// Remove unnecessary prefix
		line = line[len("Game "):]
		idIdx := strings.Index(line, ":")
		gameId, _ := strconv.ParseInt(line[:idIdx], 10, 64)
		for _, sample := range strings.Split(line[idIdx+2:], ";") {
			// Split by comma
			sample = strings.TrimSpace(sample)
			for _, ball := range strings.Split(sample, ",") {
				ball = strings.TrimSpace(ball)
				if len(ball) == 0 {
					continue
				}
				ballIdx := strings.Index(ball, " ")
				ballCount, _ := strconv.ParseInt(ball[:ballIdx], 10, 64)
				ballColor := ball[ballIdx+1:]
				// Update minGameBag
				minGameBag[ballColor] = max(minGameBag[ballColor], ballCount)
				if ballCount > bagLoad[ballColor] && gamePossible {
					impossibleGameIdSum += int(gameId)
					// Don't double count
					gamePossible = false
				}
			}
		}
		gamePower := 1
		for _, ballColor := range []string{"red", "green", "blue"} {
			gamePower *= int(minGameBag[ballColor])
		}
		gamePowerSum += gamePower
		if gamePossible {
			possibleGameIdSum += int(gameId)
		}
	}
	fmt.Println("Problem 2A:", possibleGameIdSum)
	fmt.Println("Problem 2B:", gamePowerSum)
}

func Problem1() {
	calLines := helpers.LoadInputLines(2023, 1, true)
	calValue := 0
	nameCalValue := 0
	digitNames := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range calLines {
		fc := -1
		fc2 := -1
		lc := -1
		lc2 := -1
		for i, c := range line {
			subL := line[i:]
			for d, name := range digitNames {
				if strings.HasPrefix(subL, name) {
					lc2 = d
					if fc2 == -1 {
						fc2 = d
					}
					break
				}
			}
			if unicode.IsDigit(c) {
				lc, _ = strconv.Atoi(string(c))
				if fc == -1 {
					fc = lc
				}
				if fc2 == -1 {
					fc2 = lc
				}
				lc2 = lc
			}
		}
		calValue += fc*10 + lc
		nameCalValue += fc2*10 + lc2
	}
	fmt.Println("Problem 1A:", calValue)
	fmt.Println("Problem 1B:", nameCalValue)
}
