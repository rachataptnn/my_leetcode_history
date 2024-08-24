// https://leetcode.com/problems/find-the-closest-palindrome/description/?envType=daily-question&envId=2024-08-24

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "807045053224792883"
	fmt.Println(nearestPalindromic(n))
}

func nearestPalindromic(n string) string {
	padding := 1
	nInt, _ := strconv.Atoi(n)
	for {
		leftNum := strconv.Itoa(nInt - padding)
		rightNum := strconv.Itoa(nInt + padding)
		if isPalindrome(leftNum) {
			return leftNum
		}
		if isPalindrome(rightNum) {
			return rightNum
		}
		padding++
	}
}

func isPalindrome(str string) bool {
	indexFirst := 0
	indexLast := len(str) - 1

	for indexFirst < indexLast {
		if str[indexFirst] != str[indexLast] {
			return false
		}
		indexFirst++
		indexLast--
	}

	return true
}
