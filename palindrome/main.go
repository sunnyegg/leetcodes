package palindrome

import (
	"math"
	"strconv"
)

func Calc(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	halfLen := int(math.Floor(float64(len(str)) / 2))

	for i := 0; i < halfLen; i++ {
		first := i
		second := len(str) - 1 - i

		if string(str[first]) != string(str[second]) {
			return false
		}
	}

	return true
}
