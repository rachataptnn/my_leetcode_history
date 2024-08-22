// https://leetcode.com/problems/reverse-integer/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	res := reverseInt(x)
	return res
}

func reverseInt(num int) int {
	str := strconv.Itoa(num)
	reversedStr := reverseString(str)
	reversedInt, _ := strconv.Atoi(reversedStr)

	return reversedInt
}

func reverseString(s string) string {
	runes := []rune(s)
	isNegative := false

	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedStr := string(runes)

	if isNegative {
		reversedStr = "-" + reversedStr
		return reversedStr[:len(reversedStr)-1]
	}

	return reversedStr
}
