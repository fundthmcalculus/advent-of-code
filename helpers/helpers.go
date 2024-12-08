package helpers

import (
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func LoadInputLines(year int, day int, test bool) []string {
	l := strings.Split(LoadInput(year, day, test), "\n")
	for i := 0; i < len(l); i++ {
		l[i] = strings.TrimSpace(l[i])
	}
	return l
}

func LoadInput(year int, day int, test bool) string {
	testSuffix := ""
	if test {
		testSuffix = "-test"
	}
	b, err := os.ReadFile(path.Join(".", fmt.Sprintf("aoc%d", year), fmt.Sprintf("p%d%s.txt", day, testSuffix)))

	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func AllSame[S string](x []S, y S) bool {
	for _, q := range x {
		if q != y {
			return false
		}
	}
	return true
}

func ToInt(x []string) []int {
	var y []int
	for i := 0; i < len(x); i++ {
		if x[i] == "" {
			continue
		}
		xi, _ := strconv.ParseInt(x[i], 10, 64)
		y = append(y, int(xi))
	}
	return y
}

func ToInt64(x ...string) []int64 {
	var y []int64
	for i := 0; i < len(x); i++ {
		if x[i] == "" {
			continue
		}
		xi, _ := strconv.ParseInt(x[i], 10, 64)
		y = append(y, xi)
	}
	return y
}

func Max[E Numeric](x ...E) E {
	mx := E(0)
	for _, es := range x {
		if es > mx {
			mx = es
		}
	}
	return mx
}

func Add[S ~[]E, E Numeric](x E, y S) S {
	var z S
	for _, y1 := range y {
		z = append(z, x+y1)
	}
	return z
}

func Range[S Numeric](x S, y S) []S {
	var z []S
	for i := x; i <= y; i++ {
		z = append(z, i)
	}
	return z
}

func Pow[S Numeric](x S, y S) S {
	return S(math.Pow(float64(x), float64(y)))
}

func Sum[S Numeric](x ...S) S {
	var z S
	for _, x1 := range x {
		z += x1
	}
	return z
}

type Numeric interface {
	// TODO - Other types
	int | float64 | int32 | int64
}

func Intersect[S ~[]E, E Numeric](x S, y S) S {
	//compFunc := func(a, b E) int {
	//	if a < b {
	//		return -1
	//	}
	//	if a == b {
	//		return 0
	//	}
	//	return 1
	//}
	//slices.SortFunc(x, compFunc)
	//slices.SortFunc(y, compFunc)
	var z S
	for _, x1 := range x {
		for _, y1 := range y {
			if x1 == y1 {
				z = append(z, x1)
			}
		}
	}
	return z
}

func SumSquare[S ~[]E, E Numeric](x S) float64 {
	sS := 0.0
	for _, x1 := range x {
		sS += float64(x1) * float64(x1)
	}
	return sS
}

func InSlice[S ~[]E, E Numeric](x S, y E) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == y {
			return true
		}
	}
	return false
}

func Magnitude[S ~[]E, E Numeric](x S) float64 {
	return math.Sqrt(SumSquare(x))
}

func Atoi(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
