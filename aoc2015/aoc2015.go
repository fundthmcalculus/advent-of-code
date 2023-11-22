package aoc2015

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

const aoc2015 = "aoc2015"

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
	data := loadInput(2)
	packages := strings.Split(data, "\n")
	totalWrap := 0
	totalRibbon := 0
	for _, p := range packages {
		dims := toInt(strings.Split(strings.TrimSpace(p), "x"))
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
