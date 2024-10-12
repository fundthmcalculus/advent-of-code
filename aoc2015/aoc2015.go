package aoc2015

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/fundthmcalculus/advent-of-code/helpers"
	_ "github.com/fundthmcalculus/advent-of-code/helpers"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const aoc2015 = "aoc2015"

func Problem1() {
	// Load p1.txt
	lispData := helpers.LoadInput(2015, 1, false)
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
	packages := helpers.LoadInputLines(2015, 2, false)
	totalWrap := 0
	totalRibbon := 0
	for _, p := range packages {
		dims := helpers.ToInt(strings.Split(p, "x"))
		slices.Sort(dims)
		wrap := 3*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
		totalWrap += wrap
		totalRibbon += 2*(dims[0]+dims[1]) + dims[0]*dims[1]*dims[2]
	}
	fmt.Println("Problem 2A Answer:", totalWrap)
	fmt.Println("Problem 2B Answer:", totalRibbon)
}

func Problem3() {
	directions := helpers.LoadInput(2015, 3, false)
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
	words := helpers.LoadInputLines(2015, 5, false)
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
	instructions := helpers.LoadInputLines(2015, 6, false)
	re := regexp.MustCompile(`(?m)(?P<step>turn on|turn off|toggle)\s*(?P<x1>\d+),(?P<y1>\d+)\s*\w*\s*(?P<x2>\d+),(?P<y2>\d+)`)
	// Initializing to everything off
	var lights [1000][1000]bool
	var nordicLights [1000][1000]int
	for _, instruction := range instructions {
		matchSteps := re.FindStringSubmatch(instruction)
		coordinates := helpers.ToInt(matchSteps[2:])
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
	circuitLines := helpers.LoadInputLines(2015, 7, false)
	circuits := evaluateCircuit(circuitLines)
	aVal := circuits["a"]
	// Crude hack to override "b"
	circuitLines = helpers.LoadInputLines(2015, 7, false)
	circuitLines[89] = fmt.Sprintf("%d -> b", aVal)
	circuits = evaluateCircuit(circuitLines)
	fmt.Println("Problem 7A Answer:", aVal)
	fmt.Println("Problem 7B Answer:", circuits["a"])
}

func Problem8() {
	stringLines := helpers.LoadInputLines(2015, 8, false)
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
	distValues := helpers.LoadInputLines(2015, 9, false)
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
	// Run the test operations
	num0 := "1"
	fmt.Println("Problem 10 Test:")
	for i := 0; i <= 4; i++ {
		num0 = lookAndSay(num0)
		fmt.Println(num0)
	}
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

func Problem11() {
	fmt.Println("Problem 11 Tests:")
	problem11Requirements(str2uint8slice("hijklmmn"))
	problem11Requirements(str2uint8slice("abbceffg"))
	problem11Requirements(str2uint8slice("abbcegjk"))

	fmt.Println("Test Password:", findValidPassword("abcdefgh"))
	passwordA := findValidPassword("hepxcrrq")
	fmt.Println("Problem 11A Answer:", passwordA)
	fmt.Println("Problem 11B Answer:", findValidPassword(passwordA))
}

func Problem12() {
	// Load the JSON file
	jsonData := helpers.LoadInput(2015, 12, false)
	var sum int
	// Iterate one character at a time
	for i := 0; i < len(jsonData); i++ {
		// Look for a number
		if jsonData[i] >= '0' && jsonData[i] <= '9' {
			// Find the end of the number
			j := i
			for j < len(jsonData) && jsonData[j] >= '0' && jsonData[j] <= '9' {
				j++
			}
			// Check if the previous character was a negative sign
			if i > 0 && jsonData[i-1] == '-' {
				i--
			}
			// Convert to integer
			num, _ := strconv.Atoi(jsonData[i:j])
			sum += num
			i = j
		}
	}
	fmt.Println("Problem 12A Answer:", sum)

	// Parse into a map, ignore red
	var sum2 int
	var jsonDataMap map[string]interface{}
	json.Unmarshal([]byte(jsonData), &jsonDataMap)
	// Recursively sum the values
	sum2 = sumJSON(jsonDataMap)
	fmt.Println("Problem 12B Answer:", sum2)
}

func sumJSON(jsonDataMap map[string]interface{}) int {
	var sum2 int
	for k, v := range jsonDataMap {
		if k == "red" {
			return 0
		}
		switch v.(type) {
		case string:
			if v.(string) == "red" {
				return 0
			}
		case float64:
			sum2 += (int)(v.(float64))
		case map[string]interface{}:
			sum2 += sumJSON(v.(map[string]interface{}))
		case []interface{}:
			sum2 += sumJSONArray(v.([]interface{}))
		}
	}
	return sum2
}

func sumJSONArray(jsonDataArray []interface{}) int {
	var sum2 int
	for _, v := range jsonDataArray {
		switch v.(type) {
		case string:
			// Ignore
		case float64:
			sum2 += (int)(v.(float64))
		case map[string]interface{}:
			sum2 += sumJSON(v.(map[string]interface{}))
		case []interface{}:
			sum2 += sumJSONArray(v.([]interface{}))
		}
	}
	return sum2
}

func findValidPassword(password string) string {
	testPassword := str2uint8slice(password)
	// Find the next valid password
	if problem11Requirements(testPassword) {
		testPassword = indexPassword(testPassword)
	}
	for !problem11Requirements(testPassword) {
		testPassword = indexPassword(testPassword)
	}
	return uint8slice2str(testPassword)
}

func uint8slice2str(s []uint8) string {
	var s2 []rune
	for _, c := range s {
		s2 = append(s2, rune(c))
	}
	return string(s2)
}

func str2uint8slice(s string) []uint8 {
	var s2 []uint8
	for _, c := range s {
		s2 = append(s2, uint8(c))
	}
	return s2
}

func indexPassword(password []uint8) []uint8 {
	// Find the last character that is not 'z'
	password[len(password)-1]++
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] > 'z' {
			password[i] = 'a'
			if i > 0 {
				password[i-1]++
			}
		} else {
			break
		}
	}
	return password
}

func problem11Requirements(password []uint8) bool {
	return problem11Requirement1(password) && problem11Requirement2(password) && problem11Requirement3(password)
}

func problem11Requirement3(password []uint8) bool {
	// Passwords must contain at least two different, non-overlapping pairs of letters
	pairCount := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairCount++
			// Skip the next one so we don't get "aaa" as "aa" and "aa"
			i++
		}
	}
	return pairCount >= 2
}

func problem11Requirement2(password []uint8) bool {
	// Passwords may not contain the letters i, o, or l
	// Check for forbidden letters
	for _, c := range password {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
	}
	return true
}

func problem11Requirement1(password []uint8) bool {
	// Check for an increasing straight of 3
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
			return true
		}
	}
	return false
}

func lookAndSay2(num []uint8) []uint8 {
	// Loop through string, replacing run of digits with count
	var newNum []uint8
	for idx := 0; idx < len(num); {
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
		// Convert chrCnt into sequence of digits
		for divisor := int(math.Log10(float64(chrCnt))); divisor >= 0; divisor-- {
			digit := chrCnt / int(math.Pow10(divisor))
			newNum = append(newNum, uint8(digit))
			chrCnt -= digit * int(math.Pow10(divisor))
		}
		newNum = append(newNum, cchr)
	}
	return newNum
}

func lookAndSay(num string) string {
	// Convert string into slice of integer
	numSlice := make([]uint8, len(num))
	for i := 0; i < len(num); i++ {
		numSlice[i] = num[i] - '0'
	}
	newNum := lookAndSay2(numSlice)
	// Convert back to string, index by '0'
	for i := 0; i < len(newNum); i++ {
		newNum[i] += '0'
	}
	return string(newNum)
}

func findShortestRoute(distMat [][]int, usedCities []int, fromCity int) int {
	minDist := math.MaxInt
	// Run through unused cities
	for toCity := 0; toCity < len(distMat); toCity++ {
		// Skip used cities
		if !helpers.InSlice(usedCities, toCity) {
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
		A := helpers.Magnitude(a[:])
		B := helpers.Magnitude(b[:])
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
