// https://leetcode.com/problems/palindrome-number/description/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

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
