// https: //leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/description/?envType=daily-question&envId=2024-08-03

package main

import (
	"fmt"
	"sort"
)

func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	fmt.Println(canBeEqual(target, arr))
}

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, v := range target {
		if v != arr[i] {
			return false
		}
	}

	return true
}
