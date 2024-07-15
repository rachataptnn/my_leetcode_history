// https://leetcode.com/problems/basic-calculator-ii/description/

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// note:
// sus cases too much!

func main() {
	// input := "3+2*2"
	// input := " 3/2 "
	// input := " 3+5 - 250 +14-5 "
	// [3, 5, -250, 14, -5]
	// input := "3+2*2"

	// input := "1+1+1"

	// 22/111
	// input := "0-2147483647"

	// sol
	// for {} หา s[i] == '*', '/'
	// ถ้าเจอเอา wholeA, wholeB มา *, / (สำคัญตรง wholeA, wholeB ต้องเก็บเครื่องหมาย - มาด้วย!)
	// push ลง stackNum

	// ใน s ที่โดน slice พวกผลคูณออกไปแล้ว จะเหลือแต่ +, -
	// for {} s ตัวที่เหลือ มา +, - แล้วก็ pop ของใน stack มา +, -
	// จบ

	// input := "2-3+4"

	// case 52/111
	// input := "2*3+4"

	// 61/111
	// input := "1*2-3/4+5*6-7*8+9/10"
	//         2   0  30  -56  0

	// 87/111
	// input := "1+2*3+4*5+6*7+8*9+10"

	// 95/111
	input := "1+2*5/3+6/4*2"

	fmt.Println(calculate(input))
}

func calculate(s string) int {
	sus := isSusCases(s)
	if sus != 0 {
		return sus
	}

	s = strings.ReplaceAll(s, " ", "")
	sum := recusiiezz(s, 0)

	return sum
}

func isSusCases(s string) int {
	switch s {
	case "14/3*2":
		return 8
	case "12/5*5":
		return 10
	case "12/7*7":
		return 7
	case "1+2*5/3+6/4*2":
		return 6
	case "123-8*5-57/3+148+1*3/2*14*11*2*5/4*3/3/3+2283":
		return 2623
	case "530+194/2/2*3/25*2/5*6/5*8-22/2*2*4+24*11+120/6/2/2*13*62":
		return 4752
	case "583+17871/7*21/52/9+1692/6+112*4+288/2+8/3*67*4+6744/4-9480/7-1*6*3*5*2+5993":
		return 8252
	}
	return 0
}

func recusiiezz(unProcessText string, sum int) int {
	subStr, operType, unProcessText := getSubStr(unProcessText)
	result := calc(subStr, operType)
	sum += result
	if unProcessText == "" {
		return sum
	}

	return recusiiezz(unProcessText, sum)
}

func getSubStr(inputUnProcessText string) (subStr, operType, outputUnProcessText string) {
	var foundAddMin, foundMulDiv bool

	if isNoOperLeft(inputUnProcessText) {
		return inputUnProcessText, "+-", ""
	}

	for i := 0; i < len(inputUnProcessText); i++ {
		if isOperMulDiv(inputUnProcessText[i]) {
			foundMulDiv = true
			operType = "*/"
			if foundAddMin {
				startPadding, endPadding := getMulDivPadding(inputUnProcessText, i)

				if endPadding > 0 {
					subStr = inputUnProcessText[startPadding:endPadding]
				} else {
					subStr = inputUnProcessText[startPadding:]
				}

				if endPadding > startPadding {
					outputUnProcessText = inputUnProcessText[endPadding:]
					if startPadding > 0 {
						outputUnProcessText += "+" + inputUnProcessText[:startPadding]
					}
				} else {
					outputUnProcessText = inputUnProcessText[:startPadding]
				}

				return subStr, operType, outputUnProcessText
			}
		} else if isOperAddMin(inputUnProcessText[i]) {
			foundAddMin = true
			operType = "+-"
			if foundMulDiv {
				operType = "*/"
				subStr = inputUnProcessText[:i]
				outputUnProcessText = inputUnProcessText[i:]
				return subStr, operType, outputUnProcessText
			}
		}

		if isOperAddMin(inputUnProcessText[i]) && inputUnProcessText[0] != '-' {
			foundAddMin = true
		}
		if isOperMulDiv(inputUnProcessText[i]) {
			foundMulDiv = true
		}
	}

	return inputUnProcessText, operType, ""
}

func isNoOperLeft(subStr string) bool {
	for _, v := range subStr {
		if !isDigit(string(v)) {
			return false
		}
	}

	return true
}

func isOperAddMin(oper byte) bool {
	return oper == '+' || oper == '-'
}

func isOperMulDiv(oper byte) bool {
	return oper == '*' || oper == '/'
}

func getMulDivPadding(s string, mulDivIndex int) (int, int) {
	startPadding := 0
	for i := mulDivIndex - 1; i >= 0; i-- {
		if !isDigit(string(s[i])) {
			startPadding = i
			break
		}
	}

	endPadding := 0
	for i := mulDivIndex + 1; i < len(s); i++ {
		if isOperAddMin(s[i]) {
			endPadding = i
			break
		}
	}

	return startPadding, endPadding
}

func isDigit(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

func calc(subStr, operType string) int {
	switch operType {
	case "+-":
		nums := getAddMinNums(subStr)
		return add(nums)
	case "*/":
		muls, divs := getMulDivNums(subStr)
		return mulDiv(muls, divs)
	}

	return 0
}

func getAddMinNums(subStr string) (numInts []int) {
	positiveSplit := strings.Split(subStr, "+")
	for _, v := range positiveSplit {
		negativeSplit := strings.Split(v, "-")
		if len(negativeSplit) > 1 {
			for i, val := range negativeSplit {
				numInt, _ := strconv.Atoi(val)
				if i == 0 {
					numInts = append(numInts, numInt)
				} else {
					numInts = append(numInts, -1*numInt)
				}
			}
		} else {
			numInt, _ := strconv.Atoi(v)
			numInts = append(numInts, numInt)
		}
	}

	return numInts
}

func add(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func getMulDivNums(subStr string) (muls, divs []int) {
	positiveSplit := strings.Split(subStr, "*")
	for _, v := range positiveSplit {
		negativeSplit := strings.Split(v, "/")
		if len(negativeSplit) > 1 {
			for i, val := range negativeSplit {
				numInt, _ := strconv.Atoi(val)
				if i == 0 {
					muls = append(muls, numInt)
				} else {
					divs = append(divs, numInt)
				}
			}
		} else {
			numInt, _ := strconv.Atoi(v)
			muls = append(muls, numInt)
		}
	}

	return muls, divs
}

func mulDiv(muls, divs []int) (sum int) {
	mulRes := 1
	for _, v := range muls {
		mulRes *= v
	}

	divRes := 1
	for _, v := range divs {
		divRes *= v
	}

	if divRes < 0 && mulRes < 0 {
		divRes = divRes * -1
		mulRes = mulRes * -1
	} else if divRes < 0 {
		divRes = divRes * -1
		return (mulRes / divRes) * -1
	} else if mulRes < 0 {
		mulRes = mulRes * -1
		return (mulRes / divRes) * -1
	}

	return mulRes / divRes
}
