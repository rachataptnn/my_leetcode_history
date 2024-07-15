// https://leetcode.com/problems/two-sum/description/

package main

import "fmt"

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9

	nums := []int{3, 3}
	target := 6

	// nums := []int{3, 2, 3}
	// target := 6

	// nums := []int{2, 5, 5, 11}
	// target := 10
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}
