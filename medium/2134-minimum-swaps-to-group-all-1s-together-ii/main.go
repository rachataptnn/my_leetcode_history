// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii/description/?envType=daily-question&envId=2024-08-02

package main

import (
	"fmt"
)

func main() {

	// nums := []int{0, 1, 0, 1, 1, 0, 0}
	// nums := []int{0, 1, 1, 1, 0, 0, 1, 1, 0}
	// nums := []int{1, 1, 0, 0, 1}
	// nums := []int{1}
	nums := []int{1, 1, 1, 0, 0, 1, 0, 1, 1, 0}

	fmt.Println(minSwaps(nums))
}
func minSwaps(nums []int) int {
	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}

	swaps := 0
	for i := len(nums) - ones; i < len(nums); i++ {
		if nums[i] == 0 {
			swaps++
		}
	}

	min_swaps := swaps
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			swaps++
		}

		last := i - ones

		if last < 0 {
			last += len(nums)
		}

		if nums[last] == 0 {
			swaps--
			if swaps < min_swaps {
				min_swaps = swaps
			}
		}
	}
	return min_swaps
}
