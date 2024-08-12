// https://leetcode.com/problems/first-missing-positive/

package main

import "fmt"

func main() {
	// nums := []int{1, 2, 0}
	// nums := []int{3, 4, -1, 1}

	nums := []int{7, 8, 9, 11, 12}

	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	foundMap := make(map[int]bool)
	max := 0
	for _, num := range nums {
		if num > 0 {
			foundMap[num] = true
		}
		if num > max {
			max = num
		}
	}

	for i := 1; i < max; i++ {
		_, ok := foundMap[i]
		if !ok {
			return i
		}
	}

	return max + 1
}
