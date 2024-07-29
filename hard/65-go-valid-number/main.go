// https://leetcode.com/problems/valid-number/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// "0"
	s := "."

	// 967 / 1494
	// s := ".1"

	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	result := buildInCheckNum(s)
	if result {
		return true
	}

	if isOneDecimalPoint(s) {
		if isReallyDecimalRight(s) {
			return true
		}
	}

	return false
}

// this func only make 967 / 1494 testcases passed
func buildInCheckNum(s string) bool {
	_, err := strconv.Atoi(s)
	res := err == nil
	return res
}

func isOneDecimalPoint(s string) bool {
	count := strings.Count(s, ".")
	return count == 1
}

func isReallyDecimalRight(s string) bool {
	split := strings.Split(s, ".")
	if buildInCheckNum(split[0]) {
		return true
	}

	if buildInCheckNum(split[1]) {
		return true
	}

	return false
}
