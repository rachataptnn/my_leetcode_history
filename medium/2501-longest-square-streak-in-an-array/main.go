// https://leetcode.com/problems/longest-square-streak-in-an-array/description/?envType=daily-question&envId=2024-10-28\

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 3, 6, 16, 8, 2}
	res := longestSquareStreak(arr)
	fmt.Println(res)
}

func longestSquareStreak(nums []int) int {
	longest := -1

	allSubsets := allSubsets(nums)
	// fmt.Println(allSubsets)

	for _, subArr := range allSubsets {
		sort.Ints(subArr)

		if isSquareStreak(subArr) {
			if len(subArr) > longest {
				longest = len(subArr)
			}
		}
	}

	return longest
}

func isSquareStreak(subArr []int) bool {
	for i := 0; i < len(subArr)-1; i++ {
		if subArr[i]*subArr[i] != subArr[i+1] {
			return false
		}
	}

	return true
}

func generateSubsets(arr []int, start int, subset []int, result *[][]int) {
	if len(subset) >= 2 {
		temp := make([]int, len(subset))
		copy(temp, subset)
		*result = append(*result, temp)
	}

	for i := start; i < len(arr); i++ {
		subset = append(subset, arr[i])           // Include arr[i] in the subset
		generateSubsets(arr, i+1, subset, result) // Recurse with remaining elements
		subset = subset[:len(subset)-1]           // Backtrack to remove arr[i]
	}
}

func allSubsets(arr []int) [][]int {
	var result [][]int
	generateSubsets(arr, 0, []int{}, &result)
	return result
}
