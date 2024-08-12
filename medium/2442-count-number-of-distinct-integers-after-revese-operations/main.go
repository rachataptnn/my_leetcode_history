// https://leetcode.com/problems/count-number-of-distinct-integers-after-reverse-operations/description/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nums := []int{1, 13, 10, 12, 31}
	nums := []int{2, 2, 2}
	fmt.Println(countDistinctIntegers(nums))
}

func countDistinctIntegers(nums []int) int {
	revNums := reverseEachNums(nums)
	nums = append(nums, revNums...)

	distinctCnt := 0
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
		if freqMap[num] < 2 {
			distinctCnt++
		}
	}

	return distinctCnt
}

func reverseEachNums(nums []int) (revNums []int) {
	for _, num := range nums {
		numStr := strconv.Itoa(num)
		revStr := ""
		for i := len(numStr) - 1; i >= 0; i-- {
			revStr += string(numStr[i])
		}
		revNum, _ := strconv.Atoi(revStr)
		revNums = append(revNums, revNum)

	}

	return revNums
}
