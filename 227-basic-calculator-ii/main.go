// https://leetcode.com/problems/basic-calculator-ii/description/

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// input := " 3/2 "
	// input := " 3+5 / 2 "
	// input := "3+2*2"

	// input := "1+1+1"

	// 22/111
	// input := "0-2147483647"

	input := "2-3+4"

	fmt.Println(calculate(input))
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")

	if !isMulDivGone(s) {
		s = removeMulDiv(s)
	}

	if !isAddMinGone(s) {
		s = removeAddMin(s)
	}

	res, _ := strconv.Atoi(s)

	return res
}

// mul
func removeMulDiv(s string) string {
	i := 1
	res := 0
	for {
		if isAddMin(s[i]) {
			i += 2
		}

		if isMulDiv(s[i]) {
			switch s[i] {
			case '*':
				a, b, padding := getWholeAB(s, i)
				res = mul(a, b)

				headPart := s[:i-1]
				calculated := strconv.Itoa(res)
				tailPart := s[i+2+padding:]
				s = headPart + calculated + tailPart

			case '/':
				a, b, padding := getWholeAB(s, i)
				res = divide(a, b)

				headPart := s[:i-1]
				calculated := strconv.Itoa(res)
				tailPart := s[i+2+padding:]
				s = headPart + calculated + tailPart
			}

		}

		if isMulDivGone(s) {
			break
		}

		i = 1
	}

	return s
}

func mul(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func isMulDiv(token byte) bool {
	if token == '*' || token == '/' {
		return true
	}

	return false
}

func isMulDivGone(s string) bool {
	for _, v := range s {
		if v == '*' || v == '/' {
			return false
		}
	}

	return true
}

// add
func removeAddMin(s string) string {
	i := 1
	res := 0
	for {
		if isMulDiv(s[i]) {
			i += 2
		}

		if isResultMinus(s) {
			break
		}

		if isAMinus(string(s[0])) {
			res, padding := getWholeMinusAAndB(s)

			calculated := strconv.Itoa(res)
			tailPart := s[padding:]
			s = calculated + tailPart
			fmt.Println("")
		}

		if isAddMinGone(s) {
			break
		}

		if isAddMin(s[i]) {
			switch s[i] {
			case '+':
				a, b, padding := getWholeAB(s, i)
				res = add(a, b)

				headPart := s[:i-1]
				calculated := strconv.Itoa(res)
				tailPart := s[i+2+padding:]
				s = headPart + calculated + tailPart

			case '-':
				a, b, padding := getWholeAB(s, i)
				res = minus(a, b)

				headPart := s[:i-1]
				calculated := strconv.Itoa(res)
				tailPart := s[i+2+padding:]
				s = headPart + calculated + tailPart
			}

		}

		i = 1
	}

	return s
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func isAddMin(token byte) bool {
	if token == '+' || token == '-' {
		return true
	}

	return false
}

func isAddMinGone(s string) bool {
	for _, v := range s {
		if v == '+' || v == '-' {
			return false
		}
	}

	return true
}

func getWholeAB(s string, i int) (int, int, int) {
	centerIndex := i

	firstToken := 0
	for i := centerIndex; i >= 0; i-- {
		if !isDigit(string(s[i])) {
			firstToken = i
			break
		}
	}

	lastToken := len(s)
	for i := centerIndex; i < len(s); i++ {
		if !isDigit(string(s[i])) {
			lastToken = i
			break
		}
	}

	a := s[firstToken-1 : centerIndex]
	aInt, _ := strconv.Atoi(a)

	b := s[centerIndex+1 : lastToken+2]
	bInt, _ := strconv.Atoi(b)

	return aInt, bInt, len(b) - 1
}

func isDigit(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

func isAMinus(input string) bool {
	if input == "-" {
		return true
	}
	return false
}

func getWholeMinusAAndB(input string) (int, int) {
	lastIndexA := 0
	for i := 1; i < len(input); i++ {
		if !isDigit(string(input[i])) {
			lastIndexA = i
			break
		}
	}
	a := input[:lastIndexA]

	lastIndexB := len(input)
	for i := lastIndexA + 1; i < len(input); i++ {
		if !isDigit(string(input[i])) {
			lastIndexB = i
			break
		}
	}
	b := input[lastIndexA+1 : lastIndexB]

	padding := len(a) + 1 + len(b)

	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)

	res := 0
	switch input[lastIndexA : lastIndexB-1] {
	case "+":
		res = aInt + bInt
	case "-":
		res = aInt - bInt
	}

	return res, padding
}

func isResultMinus(s string) bool {
	minusCnt := 0
	addCnt := 0
	for _, v := range s {
		if v == '-' {
			minusCnt++
		} else if v == '+' {
			addCnt++
		}

	}

	if minusCnt == 1 && addCnt == 0 {
		return true
	}

	return false
}
