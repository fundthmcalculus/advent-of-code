package aoc2015

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const aoc2015 = "aoc2015"

func loadInputLines(day int) []string {
	l := strings.Split(loadInput(day), "\n")
	for i := 0; i < len(l); i++ {
		l[i] = strings.TrimSpace(l[i])
	}
	return l
}

func loadTestInputLines(day int) []string {
	l := strings.Split(loadTestInput(day), "\n")
	for i := 0; i < len(l); i++ {
		l[i] = strings.TrimSpace(l[i])
	}
	return l
}

func loadInput(day int) string {
	b, err := os.ReadFile(path.Join(".", aoc2015, fmt.Sprintf("p%d.txt", day)))

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func loadTestInput(day int) string {
	b, err := os.ReadFile(path.Join(".", aoc2015, fmt.Sprintf("p%d-test.txt", day)))

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func toInt(x []string) []int {
	var y []int
	for i := 0; i < len(x); i++ {
		xi, _ := strconv.ParseInt(x[i], 10, 32)
		y = append(y, int(xi))
	}
	return y
}

type numeric interface {
	// TODO - Other types
	int | float64 | int32 | int64
}

func sumSquare[S ~[]E, E numeric](x S) float64 {
	sS := 0.0
	for _, x1 := range x {
		sS += float64(x1) * float64(x1)
	}
	return sS
}

func inSlice[S ~[]E, E numeric](x S, y E) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == y {
			return true
		}
	}
	return false
}

func magnitude[S ~[]E, E numeric](x S) float64 {
	return math.Sqrt(sumSquare(x))
}

func Problem1() {
	// Load p1.txt
	lispData := loadInput(1)
	curFloor := 0
	basementBegin := 0
	for i := 0; i < len(lispData); i++ {
		if lispData[i] == '(' {
			curFloor++
		} else if lispData[i] == ')' {
			curFloor--
		}
		if curFloor < 0 && basementBegin == 0 {
			basementBegin = i + 1
		}
	}
	fmt.Println("Problem 1A Answer:", curFloor)
	fmt.Println("Problem 1B Answer:", basementBegin)
}

func Problem2() {
	packages := loadInputLines(2)
	totalWrap := 0
	totalRibbon := 0
	for _, p := range packages {
		dims := toInt(strings.Split(p, "x"))
		slices.Sort(dims)
		wrap := 3*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
		totalWrap += wrap
		totalRibbon += 2*(dims[0]+dims[1]) + dims[0]*dims[1]*dims[2]
	}
	fmt.Println("Problem 2A Answer:", totalWrap)
	fmt.Println("Problem 2B Answer:", totalRibbon)
}

func Problem3() {
	directions := loadInput(3)
	houses := santaRoute(directions, -1)
	uniqueHouses := gotOnePresent(houses)
	fmt.Println("Problem 3A Answer:", len(uniqueHouses))
	santaHouses := santaRoute(directions, 0)
	robotHouses := santaRoute(directions, 1)
	drunkHouses := append(santaHouses, robotHouses...)
	fmt.Println("Problem 3B Answer:", len(gotOnePresent(drunkHouses)))
}

func Problem4() {
	input := "yzbqklnj"
	key := 1
	for {
		outputData := md5.Sum([]byte(input + fmt.Sprint(key)))
		isZero := outputData[2]>>4 == 0
		for i := 0; i < 2; i++ {
			isZero = isZero && outputData[i] == 0
		}
		if isZero {
			fmt.Println("Problem 4A Answer:", key, isZero)
			break
		}
		key++
	}
	key = 1
	for {
		outputData := md5.Sum([]byte(input + fmt.Sprint(key)))
		isZero := true
		for i := 0; i < 3; i++ {
			isZero = isZero && outputData[i] == 0
		}
		if isZero {
			fmt.Println("Problem 4B Answer:", key, isZero)
			break
		}
		key++
	}
}

func Problem5() {
	words := loadInputLines(5)
	niceCount := 0
	nice2Count := 0
	for _, w := range words {
		if isNice1(w) {
			niceCount++
		}
		if isNice2(w) {
			nice2Count++
		}
	}
	fmt.Println("Problem 5A Answer:", niceCount)
	fmt.Println("Problem 5B Answer:", nice2Count)
}

func str2int(s string) (int, bool) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int(i), err == nil
}

func Problem6() {
	instructions := loadInputLines(6)
	re := regexp.MustCompile(`(?m)(?P<step>turn on|turn off|toggle)\s*(?P<x1>\d+),(?P<y1>\d+)\s*\w*\s*(?P<x2>\d+),(?P<y2>\d+)`)
	// Initializing to everything off
	var lights [1000][1000]bool
	var nordicLights [1000][1000]int
	for _, instruction := range instructions {
		matchSteps := re.FindStringSubmatch(instruction)
		coordinates := toInt(matchSteps[2:])
		x1 := coordinates[0]
		y1 := coordinates[1]
		x2 := coordinates[2]
		y2 := coordinates[3]
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				instr := matchSteps[1]
				if instr == "turn on" {
					lights[x][y] = true
					nordicLights[x][y]++
				} else if instr == "turn off" {
					lights[x][y] = false
					nordicLights[x][y]--
				} else if instr == "toggle" {
					lights[x][y] = !lights[x][y]
					nordicLights[x][y] += 2
				}
				nordicLights[x][y] = max(nordicLights[x][y], 0)
			}
		}
	}
	// Count all trues
	isLit := 0
	nordicBrightness := 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights); y++ {
			if lights[x][y] {
				isLit++
			}
			nordicBrightness += nordicLights[x][y]
		}
	}

	fmt.Println("Problem 6A Answer:", isLit)
	fmt.Println("Problem 6B Answer:", nordicBrightness)
}

func Problem7() {
	circuitLines := loadInputLines(7)
	circuits := evaluateCircuit(circuitLines)
	aVal := circuits["a"]
	// Crude hack to override "b"
	circuitLines = loadInputLines(7)
	circuitLines[89] = fmt.Sprintf("%d -> b", aVal)
	circuits = evaluateCircuit(circuitLines)
	fmt.Println("Problem 7A Answer:", aVal)
	fmt.Println("Problem 7B Answer:", circuits["a"])
}

func Problem8() {
	stringLines := loadInputLines(8)
	totalCodeCharCount := 0
	totalMemChars := 0
	totalNewEncodedCount := 0
	for _, line := range stringLines {
		totalCodeCharCount += len(line)
		// Strip the leading/trailing quotes
		totalNewEncodedCount += 6 // for the removed wrapper quotes
		if len(line) < 3 {
			continue
		}
		line = line[1 : len(line)-1]
		// Loop and count
		idx := 0
		for idx < len(line) {
			if line[idx] == '\\' {
				// Look-ahead
				laChar := line[idx+1]
				if laChar == '\\' || laChar == '"' {
					// Skip 1
					idx++
					totalNewEncodedCount += 3
				} else if laChar == 'x' {
					// Jump three
					idx += 3
					totalNewEncodedCount += 4
				} else {
					panic("Invalid escape sequence!")
				}
			} else {
			}
			totalMemChars++
			idx++
			totalNewEncodedCount++
		}
	}

	fmt.Println("Problem 8A Answer:", totalCodeCharCount-totalMemChars)
	fmt.Println("Problem 8B Answer:", totalNewEncodedCount-totalCodeCharCount)
}

func Problem9() {
	distValues := loadInputLines(9)
	N := int((math.Sqrt(float64(1+8*len(distValues))) - 1) / 2)
	// Fill out the full matrix (symmetric)
	Nc := N + 1
	var distMat = make([][]int, Nc)
	for i := 0; i < Nc; i++ {
		distMat[i] = make([]int, Nc)
	}
	// They're in order.
	mj := 0
	mi := 0
	stepRow := N + 1
	for _, dL := range distValues {
		dist, _ := str2int(dL[strings.Index(dL, "= ")+2:])
		// Principal diagonal should be empty.
		stepRow--
		if stepRow == 0 {
			mi++
			stepRow = N - mi
			mj = mi + 1
		} else {
			mj++
		}
		distMat[mi][mj] = dist
		distMat[mj][mi] = dist
	}

	// Find the shortest route - exhaustive search
	shortestRoute := math.MaxInt
	for startCity := 0; startCity < N; startCity++ {
		shortestRoute = min(shortestRoute, findShortestRoute(distMat, []int{startCity}, startCity))
	}
	// Longest route is minimum negative distance
	for i := 0; i < len(distMat); i++ {
		for j := 0; j < len(distMat); j++ {
			distMat[i][j] = -distMat[i][j]
		}
	}
	longestRoute := 0
	for startCity := 0; startCity < N; startCity++ {
		longestRoute = min(longestRoute, findShortestRoute(distMat, []int{startCity}, startCity))
	}
	fmt.Println("Problem 9A Answer:", shortestRoute)
	fmt.Println("Problem 9B Answer:", -longestRoute)
}

func Problem10() {
	// Loop through the string
	num1 := "3113322113"
	for i := 0; i < 40; i++ {
		num1 = lookAndSay(num1)
	}
	fmt.Println("Problem 10A Answer:", len(num1))
	num2 := num1
	for i := 0; i < 10; i++ {
		num2 = lookAndSay(num2)
	}
	fmt.Println("Problem 10B Answer:", len(num2))
}

func lookAndSay(num string) string {
	// Loop through string, replacing run of digits with count
	newNum := ""
	idx := 0
	for idx < len(num) {
		cchr := num[idx]
		chrCnt := 1
		// Loop forward while number matches
		for idx+chrCnt < len(num) {
			if num[idx+chrCnt] != cchr {
				break
			}
			chrCnt++
		}
		idx += chrCnt
		newNum += strconv.FormatInt(int64(chrCnt), 10) + string(cchr)
	}
	return newNum
}

func findShortestRoute(distMat [][]int, usedCities []int, fromCity int) int {
	minDist := math.MaxInt
	// Run through unused cities
	for toCity := 0; toCity < len(distMat); toCity++ {
		// Skip used cities
		if !inSlice(usedCities, toCity) {
			minDist = min(minDist, distMat[fromCity][toCity]+findShortestRoute(distMat, append(usedCities, toCity), toCity))
		}
	}
	// We've already hit everything
	if minDist == math.MaxInt {
		return 0
	}
	return minDist
}

func evaluateCircuit(circuitLines []string) map[string]uint16 {
	circuits := make(map[string]uint16)
	var wire1, wire2, op, tgt string
	// Loop forever through the input deck
	idx := 0
	for len(circuitLines) > 0 {
		line := circuitLines[idx]
		tokens := strings.Split(line, " ")

		if len(tokens) == 3 {
			// Assignment
			wire1 = tokens[0]
			tgt = tokens[2]
			wire1val, ok1 := getWireValue(circuits, wire1)
			if ok1 {
				circuits[tgt] = wire1val
				// Remove this line
				circuitLines, idx = dropCircuitLine(circuitLines, idx)
				continue
			}
		} else if len(tokens) == 4 {
			// Unary operator
			op = tokens[0]
			if op != "NOT" {
				panic("Unsupported operation")
			}
			wire1 = tokens[1]
			tgt = tokens[3]
			wire1val, ok1 := getWireValue(circuits, wire1)
			if ok1 {
				if op == "NOT" {
					circuits[tgt] = ^wire1val
				} else {
					circuits[tgt] = circuits[op]
				}
				// Remove this line
				circuitLines, idx = dropCircuitLine(circuitLines, idx)
				continue
			}
		} else if len(tokens) == 5 {
			// Binary operator
			wire1 = tokens[0]
			op = tokens[1]
			wire2 = tokens[2]
			tgt = tokens[4]
			wire1val, ok1 := getWireValue(circuits, wire1)
			wire2val, ok2 := getWireValue(circuits, wire2)
			if ok1 && ok2 {
				// Perform the operation
				if op == "AND" {
					circuits[tgt] = wire1val & wire2val
				} else if op == "OR" {
					circuits[tgt] = wire1val | wire2val
				} else if op == "LSHIFT" {
					circuits[tgt] = wire1val << wire2val
				} else if op == "RSHIFT" {
					circuits[tgt] = wire1val >> wire2val
				} else {
					panic(fmt.Sprintf("Unsupported op code %s", op))
				}
				// Remove this line
				circuitLines, idx = dropCircuitLine(circuitLines, idx)
				continue
			}
		}
		// Go to the next line, wrap back around.
		idx++
		if idx >= len(circuitLines) {
			idx = 0
		}
	}
	return circuits
}

func getWireValue(circuits map[string]uint16, wire1 string) (uint16, bool) {
	wire1val, ok1 := circuits[wire1]
	val1, isint1 := str2int(wire1)
	// Sometimes, it's a constant, that's okay
	if !ok1 && isint1 {
		wire1val = uint16(val1)
		ok1 = isint1
	}
	return wire1val, ok1
}

func dropCircuitLine(circuitLines []string, idx int) ([]string, int) {
	circuitLines = append(circuitLines[:idx], circuitLines[idx+1:]...)
	idx = 0
	return circuitLines, idx
}

func isNice2(s string) bool {
	skipRepeats := 0
	repeatPairs := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i > 0 {
			c2 := s[i-1 : i+1]
			for j := i + 1; j < len(s)-1; j++ {
				if c2 == s[j:j+2] {
					repeatPairs++
					break
				}
			}
		}
		if i > 1 && c == s[i-2] {
			skipRepeats++
		}
	}
	return repeatPairs > 0 && skipRepeats > 0
}

func isNice1(s string) bool {
	vowelCount := 0
	vowels := []uint8{'a', 'e', 'i', 'o', 'u'}
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}
	hasForbidden := false
	repeatedLetters := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if slices.Contains(vowels, c) {
			vowelCount++
		}
		if i > 0 && c == s[i-1] {
			repeatedLetters++
		}
		if i > 0 && slices.Contains(forbiddenStrings, s[i-1:i+1]) {
			hasForbidden = true
		}
	}
	return vowelCount >= 3 && repeatedLetters > 0 && !hasForbidden
}

func gotOnePresent(houses [][2]int) [][2]int {
	sortFunc := func(a, b [2]int) int {
		if a[0] == b[0] && a[1] == b[1] {
			return 0
		}
		A := magnitude(a[:])
		B := magnitude(b[:])
		if A > B {
			return 1
		} else {
			return -1
		}
	}
	compFunc := func(a, b [2]int) bool {
		return a[0] == b[0] && a[1] == b[1]
	}
	slices.SortFunc(houses, sortFunc)
	uniqueHouses := slices.CompactFunc(houses, compFunc)
	return uniqueHouses
}

func santaRoute(directions string, santaMod int) [][2]int {
	var houses [][2]int
	loc := [2]int{0, 0}
	houses = append(houses, loc)
	for i, curDir := range directions {
		// Santamod < 0 is do everything
		if santaMod >= 0 && i%2 != santaMod {
			continue
		}
		if curDir == '>' {
			loc[0]++
		} else if curDir == '<' {
			loc[0]--
		} else if curDir == '^' {
			loc[1]++
		} else if curDir == 'v' {
			loc[1]--
		}
		houses = append(houses, loc)
	}
	return houses
}
