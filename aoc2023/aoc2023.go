package aoc2023

import (
	"fmt"
	"github.com/fundthmcalculus/advent-of-code/helpers"
	"strconv"
	"strings"
	"unicode"
)

func Problem4() {
	cards := helpers.LoadInputLines(2023, 4, false)
	totalPoints := 0
	cardsWon := map[int][]int{}
	for i, card := range cards {
		// Remove unnecessary prefix
		card = card[strings.Index(card, ":")+1:]
		// Split at |
		cardHalves := strings.Split(card, "|")
		winNumbers := helpers.ToInt(strings.Split(strings.TrimSpace(cardHalves[0]), " "))
		myNumbers := helpers.ToInt(strings.Split(strings.TrimSpace(cardHalves[1]), " "))
		// Find the number of wins (intersection count)
		myWins := helpers.Intersect(winNumbers, myNumbers)
		winCnt := len(myWins)
		if winCnt > 0 {
			cardsWon[i+1] = helpers.Range(i+2, i+1+len(myWins))
			cardPoints := helpers.Pow(2, winCnt-1)
			fmt.Println("Card", i+1, "wins:", cardPoints, "points", myWins, "bonus cards", cardsWon[i+1])
			totalPoints += cardPoints
		}
	}
	cardCounts := make([]int, len(cards)+1)
	for cardId := 1; cardId <= len(cards); cardId++ {
		// Start with 1 of each card
		cardCounts[cardId] += 1
		// Now, add the cards won (which will be after this one)
		for _, wonId := range cardsWon[cardId] {
			cardCounts[wonId] += cardCounts[cardId]
		}
	}
	fmt.Println("Problem 4A:", totalPoints)
	fmt.Println("Problem 4B:", helpers.Sum(cardCounts))
}

func Problem3() {
	lines := helpers.LoadInputLines(2023, 3, false)
	partNrSum := int64(0)
	possibleGears := make(map[int]map[int][]int)
	for y := 0; y < len(lines); y++ {
		x1 := -1
		x2 := -1
		curLine := lines[y]
		for x := 0; x < len(curLine); x++ {
			// Check if this is a digit
			isDigit := unicode.IsDigit(rune(curLine[x]))
			if isDigit {
				if x1 < 0 {
					x1 = x
				}
				x2 = x
			}
			if !isDigit || x == len(curLine)-1 {
				if x1 < 0 || x2 < 0 {
					continue
				}
				// This is a number
				for iy := max(0, y-1); iy <= min(len(lines)-1, y+1); iy++ {
					for ix := max(0, x1-1); ix <= min(len(lines[iy])-1, x2+1); ix++ {
						c := lines[iy][ix]
						if !unicode.IsDigit(rune(c)) && c != '.' {
							partNr, _ := strconv.ParseInt(curLine[x1:x2+1], 10, 64)
							//fmt.Println("Part number:", partNr, string(c))

							// Gear candidates have '*'
							if c == '*' {
								//fmt.Println("Gear candidate:", ix, iy)
								// Preallocate
								if possibleGears[ix] == nil {
									possibleGears[ix] = make(map[int][]int)
								}
								possibleGears[ix][iy] = append(possibleGears[ix][iy], int(partNr))
							}

							partNrSum += partNr
							goto nextPart
						}
					}
				}
			nextPart:
				x1 = -1
				x2 = -1
			}
		}
	}
	// Gears are the ones with only 2 part Nrs
	totalRatios := 0
	for ix, possibleGearsY := range possibleGears {
		for iy, partNrs := range possibleGearsY {
			if len(partNrs) == 2 {
				gearRatio := partNrs[0] * partNrs[1]
				totalRatios += gearRatio
				//fmt.Println("Gear at", ix, iy, "with part numbers", partNrs, "Ratio:", gearRatio)
			}
		}
	}
	fmt.Println("Problem 3A:", partNrSum)
	fmt.Println("Problem 3B:", totalRatios)
}

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
