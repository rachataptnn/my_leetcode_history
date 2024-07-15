// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/

package main

import "fmt"

func main() {
	// input := []int{2, 3, 1, 2}
	// input := []int{1, 2, 10, 5, 7}
	input := []int{1, 1}
	fmt.Println(canBeIncreasing(input))
}

func canBeIncreasing(nums []int) bool {
	arr := BubbleSort(nums)
	res := BubbleSort2(arr)

	return res
}

func BubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] || numbers[j] == numbers[j+1] {
				removed := RemoveElement(numbers, j)
				return removed
			}
		}
	}

	return numbers
}

func BubbleSort2(numbers []int) bool {
	someMap := make(map[int]int)
	for _, num := range numbers {
		someMap[num]++
	}
	for _, v := range someMap {
		if v >= 2 {
			return false
		}
	}

	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				return false
			}
		}
	}

	return true
}

// RemoveElement removes an element from a slice at a specified index
func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
