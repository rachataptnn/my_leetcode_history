// https://leetcode.com/problems/find-the-closest-palindrome/description/?envType=daily-question&envId=2024-08-24

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "123"
	fmt.Println(nearestPalindromic(n))
}

func nearestPalindromic(n string) string {
	if n == "" {
		return ""
	}

	if len(n) == 1 {
		conv, _ := strconv.Atoi(n)
		dec := conv - 1
		return strconv.Itoa(dec)
	}

	halfN := len(n) / 2
	subStr := n[:halfN]
	rev := revStr(subStr)

	if len(n)%2 == 0 {
		palindrome := n[:halfN] + rev
		return palindrome
	}

	palindrome := n[:halfN+1] + rev
	return palindrome
}

func revStr(subStr string) string {
	rev := ""
	for i := len(subStr) - 1; i >= 0; i-- {
		rev += string(subStr[i])
	}

	return rev
}
