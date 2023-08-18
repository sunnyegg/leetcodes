package sqrt

import (
	"fmt"
	"math"
	"strconv"
)

// babylonian method
func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	first := float64(initialGuess(x))
	iteration := len(fmt.Sprint(x)) / 2
	var last float64

	for i := 0; i <= iteration+1; i++ {
		if i == 0 {
			last = guess(first, float64(x))
			continue
		}
		last = guess(last, float64(x))
	}

	roundDown := int(math.Floor(float64(last)))

	if roundDown < 0 {
		return 0
	} else {
		return roundDown
	}
}

func guess(x, s float64) float64 {
	return float64((x + s/x) / 2)
}

// rule of 2 and 7
// 2 or 7 for initial guess
func initialGuess(x int) int {
	length := len(fmt.Sprint(x))
	halfLen := math.Ceil(float64(length) / 2)
	zeros := fmt.Sprint(1)

	for i := 0; i < int(halfLen); i++ {
		if i == 0 {
			continue
		}
		zeros += "0"
	}

	zerosInt, _ := strconv.Atoi(zeros)

	guess2 := 2 * zerosInt
	guess7 := 7 * zerosInt

	near2 := guess2*guess2 - x
	near7 := guess7*guess7 - x

	// choose nearest to x
	if math.Abs(float64(near2)) < math.Abs(float64(near7)) {
		return guess2
	} else {
		return guess7
	}
}
