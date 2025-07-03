package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 2, 5}
	k := 2

	fmt.Println("out: ", partitionArray(nums, k))

}

func partitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	cnt := 1
	left := nums[0]
	for i := 1; i < len(nums); i++ {
		diff := abs(left, nums[i])
		if diff > k {
			left = nums[i]
			cnt++
		}
	}

	return cnt
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
