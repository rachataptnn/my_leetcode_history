// https://leetcode.com/problems/string-to-integer-atoi/

package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "20000000000000000000"
	myAtoi(input)
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	i := 0
	subNum := []byte{}

	symbol, isNegPos := isNegativePositive(s, &subNum)
	if isNegPos {
		subNum = append(subNum, symbol)
		i = 1
	}

	for {
		if i == len(s) {
			break
		}

		if isDigit(s[i]) {
			subNum = append(subNum, s[i])
		} else {
			break
		}
		i++
	}

	result, err := strconv.Atoi(string(subNum))
	if err == nil {
		if result > math.MaxInt32 {
			return math.MaxInt32
		} else if result < math.MinInt32 {
			return math.MinInt32
		}
		return int(result)
	} else if errors.Is(err, strconv.ErrRange) {
		if s[0] == '-' {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	return 0
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNegativePositive(s string, subNum *[]byte) (byte, bool) {
	if len(s) < 2 {
		return 0, false
	}

	if isDigit(s[1]) {
		switch s[0] {
		case '+':
			return '+', true
		case '-':
			return '-', true
		default:
			return 0, false
		}

	}

	return 0, false
}
