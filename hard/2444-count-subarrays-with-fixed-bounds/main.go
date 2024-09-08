// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 5, 2, 7, 5}
	minK := 1
	maxK := 5

	// nums := []int{1, 1, 1, 1}
	// minK := 1
	// maxK := 1

	fmt.Println(countSubarrays(nums, minK, maxK))
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	padding := 1
	cnt := 0

	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			subArr := nums[i:j]
			min, max := findMinMax(subArr)
			if min == minK && max == maxK {
				cnt++
			}
		}
		padding++
	}

	return int64(cnt)
}

func findMinMax(subArr []int) (min, max int) {
	min = math.MaxInt

	for _, v := range subArr {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	return min, max
}
