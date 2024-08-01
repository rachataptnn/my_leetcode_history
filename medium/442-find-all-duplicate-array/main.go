// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {
	someMap := make(map[int]int)
	res := make([]int, 0, len(nums)/2)

	for _, num := range nums {
		someMap[num]++
	}

	for k, v := range someMap {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}
