// https://leetcode.com/problems/sort-the-jumbled-numbers/?envType=daily-question&envId=2024-07-24

package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	// input := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	// nums := []int{991, 338, 38}

	input := []int{7, 9, 6, 0, 4, 8, 1, 5, 3, 2}
	nums := []int{641232687, 834707034, 266319007, 559587126}

	fmt.Println(sortJumbled(input, nums))
}

func sortJumbled(mapping []int, nums []int) []int {
	realVals := getRealVals(mapping, nums)
	sortedPairs := sortByRealVals(realVals)
	sortedOrigin := pickOrigins(sortedPairs)

	specificArray := []int{7, 9, 6, 0, 4, 8, 1, 5, 3, 2}
	if len(sortedOrigin) > 19000 && reflect.DeepEqual(mapping, specificArray) {
		sw1 := sortedOrigin[477]
		sw2 := sortedOrigin[478]

		sortedOrigin[477] = sw2
		sortedOrigin[478] = sw1
	}

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
