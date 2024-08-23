// https://leetcode.com/problems/integer-to-roman/

package main

import "fmt"

func main() {
	// num := 3749
	// num := 58
	num := 1994 // MCMXCIV
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	ts, h, t, u := getDigits(num)

	thousand := tsToRoman(ts)
	hundred := hToRoman(h)
	ten := tToRoman(t)
	unit := uToRoman(u)

	return thousand + hundred + ten + unit
}

func getDigits(num int) (int, int, int, int) {
	thousands := num / 1000 % 10
	hundreds := num / 100 % 10
	tens := num / 10 % 10
	ones := num % 10

	return thousands, hundreds, tens, ones
}

func tsToRoman(num int) string {
	if num == 0 {
		return ""
	}

	thousand := ""
	for i := 0; i < num; i++ {
		thousand += "M"
	}

	return thousand
}

func hToRoman(num int) string {
	if num == 0 {
		return ""
	}

	if num == 9 {
		return "CM"
	}

	hundred := ""
	if num > 5 {
		hundred = "D"
		for i := 5; i < num; i++ {
			hundred += "C"
		}
		return hundred
	}

	if num == 5 {
		return "D"
	}

	if num == 4 {
		return "CD"
	}

	for i := 0; i < num; i++ {
		hundred += "C"
	}

	return hundred
}

func tToRoman(num int) string {
	if num == 0 {
		return ""
	}

	if num == 9 {
		return "XC"
	}

	ten := ""
	if num > 5 {
		ten = "L"
		for i := 5; i < num; i++ {
			ten += "X"
		}
		return ten
	}

	if num == 5 {
		return "L"
	}

	if num == 4 {
		return "XL"
	}

	for i := 0; i < num; i++ {
		ten += "X"
	}

	return ten
}

func uToRoman(num int) string {
	if num == 0 {
		return ""
	}

	if num == 9 {
		return "IX"
	}

	unit := ""
	if num > 5 {
		unit = "V"
		for i := 5; i < num; i++ {
			unit += "I"
		}
		return unit
	}

	if num == 5 {
		return "V"
	}

	if num == 4 {
		return "IV"
	}

	for i := 0; i < num; i++ {
		unit += "I"
	}

	return unit
}
