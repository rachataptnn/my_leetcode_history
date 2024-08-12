package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	backtrack(&result, subset, 0, nums)

	return result
}

func backtrack(subsets *[][]int, subset []int, start int, nums []int) {
	copiedSubset := make([]int, len(subset))
	copy(copiedSubset, subset)
	*subsets = append(*subsets, copiedSubset) // prepared-subset added to subsets(result) here~

	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])      // add element to subset
		backtrack(subsets, subset, i+1, nums) // go to next element

		subset = subset[:len(subset)-1] // Backtrack: remove the last added element before going to the next iteration
	}
}
