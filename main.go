package main

import (
	"fmt"
	"leetcodes/palindrome"
	romantoint "leetcodes/romanToInt"
	"leetcodes/twosum"
)

func main() {
	// twosums
	twosumNums := []int{2, 7, 11, 15}
	twosumTarget := 9
	twosumRes := twosum.Calc(twosumNums, twosumTarget)

	fmt.Println(twosumRes)

	// palindrome
	palindromeNums := 1221
	palindromeRes := palindrome.Calc(palindromeNums)

	fmt.Println(palindromeRes)

	// roman to int
	romRes := romantoint.Calc("III")
	fmt.Println(romRes)
}
