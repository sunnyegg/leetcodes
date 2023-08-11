package main

import (
	"fmt"
	"leetcodes/twosum"
)

func main() {
	twosumNums := []int{2, 7, 11, 15}
	twosumTarget := 9
	twosumRes := twosum.Calc(twosumNums, twosumTarget)

	fmt.Println(twosumRes)
}
