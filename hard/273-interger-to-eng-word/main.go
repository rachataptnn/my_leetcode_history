package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// num := 1234567890123456789
	// num := 123
	// num := 12345

	// 290/601
	// num := 0

	// 338/601
	// num := 20

	// 428/601
	// num := 100

	// 469/601
	// num := 101

	// 514/601
	// num := 111

	//558/601
	// num := 1000

	// 583/601
	// num := 50868

	// 587/601
	num := 1000000

	fmt.Printf("\n<%+v>\n", numberToWords(num))
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	triplets := sliceIntoTriplets(num)
	words := convertTripletsToWords(triplets)
	wordsWithUnit := addUnitsToWords(words)

	return wordsWithUnit
}

type numberWord struct {
	number string
	unit   string
	word   string
}

func sliceIntoTriplets(num int) (triplets []numberWord) {
	numStr := strconv.Itoa(num)
	if len(numStr) <= 3 {
		return []numberWord{
			{
				number: numStr,
			},
		}
	}

	left := len(numStr)
	for right := len(numStr); right > 0; right -= 3 {
		left = right - 3
		if left < 0 {
			left = 0
		}

		triplets = append(triplets, numberWord{
			number: numStr[left:right],
		})
	}

	return triplets
}

func addUnitsToWords(words []string) string {
	unitList := []string{
		"",
		"Thousand",
		"Million",
		"Billion",
		"Trillion",
		"Quadrillion",
		"Quintillion",
		"Sextillion",
		"Septillion",
		"Octillion",
		"Nonillion",
	}

	res := []string{}
	for i, word := range words {
		if word == "" {
			continue
		} else {
			res = append(res, word+" "+unitList[i])
		}
	}

	realRes := ""
	for i := len(res) - 1; i >= 0; i-- {
		realRes += res[i] + " "
	}

	return strings.TrimSpace(realRes)
}

func convertTripletsToWords(tripletsWithUnit []numberWord) []string {
	digitsMap := map[rune]string{
		// '0': "Zero",
		'1': "One",
		'2': "Two",
		'3': "Three",
		'4': "Four",
		'5': "Five",
		'6': "Six",
		'7': "Seven",
		'8': "Eight",
		'9': "Nine",
	}

	teensMap := map[string]string{
		"10": "Ten",
		"11": "Eleven",
		"12": "Twelve",
		"13": "Thirteen",
		"14": "Fourteen",
		"15": "Fifteen",
		"16": "Sixteen",
		"17": "Seventeen",
		"18": "Eighteen",
		"19": "Nineteen",
	}

	tensMap := map[rune]string{
		'2': "Twenty",
		'3': "Thirty",
		'4': "Forty",
		'5': "Fifty",
		'6': "Sixty",
		'7': "Seventy",
		'8': "Eighty",
		'9': "Ninety",
	}

	realRes := []string{}
	for _, triplet := range tripletsWithUnit {
		res := []string{}
		goNext := false
		for j, char := range triplet.number {
			tmp := ""
			if goNext {
				continue
			}

			if len(triplet.number) < 3 {
				if len(triplet.number) == 1 {
					tmp += digitsMap[char]
				} else if len(triplet.number) == 2 {
					if triplet.number[0] == '1' {
						tmp += teensMap[triplet.number]
					} else {
						tmp += tensMap[rune(triplet.number[0])]

						if triplet.number[1] == 48 {
							goNext = true
						} else {
							tmp += " " + digitsMap[rune(triplet.number[1])]
						}
					}
				}
				if tmp != "" {
					res = append(res, tmp)
				}
				break
			} else {
				switch j {
				case 0:
					if digitsMap[char] != "" {
						tmp += digitsMap[char] + " " + "Hundred"
					}
				case 1:
					if char == '1' {
						ten := triplet.number[1:]
						tmp += teensMap[ten]
						goNext = true
					} else {
						tmp += tensMap[char]
					}
				case 2:
					tmp += digitsMap[char]
				}

				if tmp != "" {
					res = append(res, tmp)
				}

			}
		}

		realRes = append(realRes, strings.Join(res, " "))
	}

	return realRes
}
