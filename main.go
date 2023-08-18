package main

import (
	"fmt"
	findindexinstring "leetcodes/findIndexInString"
	longestcommonprefix "leetcodes/longestCommonPrefix"
	mergetwosortedlists "leetcodes/mergeTwoSortedLists"
	"leetcodes/palindrome"
	removeduplicates "leetcodes/removeDuplicatesFromSortedArray"
	removeelement "leetcodes/removeElement"
	romantoint "leetcodes/romanToInt"
	"leetcodes/twosum"
	validparentheses "leetcodes/validParentheses"
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

	// valid parentheses
	sortedLists := "[[[]"
	validParenthesesRes := validparentheses.IsValid(sortedLists)
	fmt.Println(validParenthesesRes)

	// merge two sorted lists
	sortedLists1 := &mergetwosortedlists.ListNode{Val: 1, Next: &mergetwosortedlists.ListNode{Val: 2, Next: &mergetwosortedlists.ListNode{Val: 4}}}
	sortedLists2 := &mergetwosortedlists.ListNode{Val: 1, Next: &mergetwosortedlists.ListNode{Val: 3, Next: &mergetwosortedlists.ListNode{Val: 4}}}
	mergedRes := mergetwosortedlists.Merge(sortedLists1, sortedLists2)
	fmt.Println(mergedRes)

	// remove duplicates in sorted array
	sortedArray := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	sortedArrayRes := removeduplicates.RemoveDuplicates(sortedArray)
	fmt.Println(sortedArrayRes)

	// remove element
	els := []int{0, 1, 2, 2, 3, 0, 4, 2}
	elsRes := removeelement.RemoveElement(els, 2)
	fmt.Println(elsRes)

	// find index string in a string
	str := "mississippi"
	strRes := findindexinstring.StrStr(str, "issip")
	fmt.Println(strRes)
}
