package main

import (
	"fmt"
	longestcommonprefix "leetcodes/longestCommonPrefix"
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
	romRes := romantoint.Calc("MCMXCIV")
	fmt.Println(romRes)

	// longest common prefix
	commonPrefix := []string{"flower", "flow", "flight"}
	commonPrefixRes := longestcommonprefix.Calc(commonPrefix)
	fmt.Println(commonPrefixRes)
}
