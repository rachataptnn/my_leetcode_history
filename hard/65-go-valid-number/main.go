// https://leetcode.com/problems/valid-number/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "3E+7"

	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	pointSeen := false
	eSeen := false
	numberSeen := false
	numberAfterE := true

	for i, ch := range s {
		if unicode.IsDigit(ch) {
			numberSeen = true
			numberAfterE = true

		} else if ch == '.' {
			if eSeen || pointSeen {
				return false // found e or . more than 1 -> false
			}
			pointSeen = true

		} else if ch == 'e' || ch == 'E' {
			if eSeen || !numberSeen {
				return false // found e before any num -> false
			}
			numberAfterE = false
			eSeen = true

		} else if ch == '-' || ch == '+' {
			if i != 0 && (s[i-1] != 'e' && s[i-1] != 'E') {
				return false // if found -e, +e -> false (i!=0 is just prevent out of range)
			}

		} else {
			return false
		}
	}

	return numberSeen && numberAfterE
}
