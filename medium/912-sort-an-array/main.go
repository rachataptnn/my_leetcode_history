// https://leetcode.com/problems/sort-an-array/description/?envType=daily-question&envId=2024-07-25

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	sort.Ints(nums)

	return nums
}
