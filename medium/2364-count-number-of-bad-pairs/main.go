package main

import "fmt"

func main() {
	nums := []int{4, 1, 3, 3}
	fmt.Println(countBadPairs(nums))
}

func countBadPairs(nums []int) int {
	n := len(nums)
	totalPairs := n * (n - 1) / 2
	diffCount := make(map[int]int)

	for i, num := range nums {
		diff := num - i
		totalPairs -= diffCount[diff]
		diffCount[diff]++
		fmt.Println("breakpoint")
	}

	return totalPairs
}
