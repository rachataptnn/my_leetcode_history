// https://leetcode.com/problems/valid-number/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := ".0e7"

	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	// s = strings.ReplaceAll(s, "0", "")

	result := isBuildInSayValid(s)
	if result {
		return true
	}

	if isDeciExpoOccur(s) {
		return false
	}

	if isSusCases(s) {
		return false
	}

	if isOneDecimalPoint(s) {
		if isReallyDecimalRight(s) {
			return true
		}
	}

	if isOneExpo(s) {
		if isReallyExpoRight(s) {
			return true
		}
	}

	return false
}

func isBuildInSayValid(s string) bool {
	_, err := strconv.Atoi(s)
	res := err == nil
	return res
}

// decimal funcs
func isOneDecimalPoint(s string) bool {
	count := strings.Count(s, ".")
	return count == 1
}

func isReallyDecimalRight(s string) bool {
	split := strings.Split(s, ".")

	if s == "+." {
		return false
	}

	if len(split[0]) > 0 {
		isFrontValid := isBuildInSayValid(split[0])
		if !isFrontValid && split[0] != "+" {
			return false
		}
	}

	if len(split[1]) > 0 {
		isBackValid := isBuildInSayValid(split[1])
		if !isBackValid {
			return false
		}

		if split[1][0] == '+' {
			return false
		}
	}

	if len(split[0]) == 0 && len(split[1]) == 0 {
		return false
	}

	return true
}

// expo funcs
func isOneExpo(s string) bool {
	count := strings.Count(s, "e")
	return count == 1
}

func isReallyExpoRight(s string) bool {
	split := strings.Split(s, "e")

	if len(split[0]) == 0 || len(split[1]) == 0 {
		return false
	}

	isFrontValid := isBuildInSayValid(split[0])
	if !isFrontValid {
		return false
	}

	isBackValid := isBuildInSayValid(split[1])
	if !isBackValid {
		return false
	}

	return true
}

// deci point & expo

func isDeciExpoOccur(s string) bool {
	if strings.Count(s, ".e") > 0 || strings.Count(s, "e.") > 0 {
		return true
	}
	return false
}

// sus case

func isSusCases(s string) bool {
	if strings.Count(s, ".-") > 0 {
		return true
	}
	return false
}
