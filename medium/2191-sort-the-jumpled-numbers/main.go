// https://leetcode.com/problems/sort-the-jumbled-numbers/?envType=daily-question&envId=2024-07-24

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	nums := []int{991, 338, 38}

	fmt.Println(sortJumbled(input, nums))
}

func sortJumbled(mapping []int, nums []int) []int {
	realVals := getRealVals(mapping, nums)
	sortedPairs := sortByRealVals(realVals)
	sortedOrigin := pickOrigins(sortedPairs)

	return sortedOrigin
}

func getRealVals(mapping []int, nums []int) (realVals [][2]int) {
	for _, num := range nums {
		realValStr := ""
		numStr := strconv.Itoa(num)
		for _, digit := range numStr {
			digitInt, _ := strconv.Atoi(string(digit))
			realDigitInt := mapping[digitInt]
			realDIgitStr := strconv.Itoa(realDigitInt)
			realValStr += realDIgitStr
		}
		realValInt, _ := strconv.Atoi(string(realValStr))
		realVals = append(realVals, [2]int{num, realValInt})
	}

	return realVals
}

func sortByRealVals(realVals [][2]int) [][2]int {
	sort.Slice(realVals, func(i, j int) bool {
		return realVals[i][1] < realVals[j][1]
	})

	return realVals
}

func pickOrigins(sortedPairs [][2]int) (origin []int) {
	for _, v := range sortedPairs {
		origin = append(origin, v[0])
	}

	return origin
}
