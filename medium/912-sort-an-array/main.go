// https://leetcode.com/problems/sort-an-array/description/?envType=daily-question&envId=2024-07-25

package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	sorted := mergeSort(nums)

	return sorted
}

func mergeSort(arr []int) []int {
	// Base case: if the array has 1 or fewer elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point
	mid := len(arr) / 2

	// Recursively sort the left and right halves
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	// Compare elements from both slices and add the smaller one to the result
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// Add any remaining elements from the left slice
	result = append(result, left[leftIndex:]...)

	// Add any remaining elements from the right slice
	result = append(result, right[rightIndex:]...)

	return result
}
