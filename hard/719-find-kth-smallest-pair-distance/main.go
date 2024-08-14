// https://leetcode.com/problems/find-k-th-smallest-pair-distance/?envType=daily-question&envId=2024-08-14

package main

import (
	"fmt"
	"sort"
)

type input struct {
	nums []int
	k    int
}

func main() {
	// nums := []int{1, 3, 1}
	// k := 1

	// nums := []int{1, 1, 1}
	// k := 2

	// nums := []int{1, 6, 1}
	// k := 3

	// nums := []int{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}
	// k := 18

	input := input{
		nums: []int{1, 2, 0, 2, 1, 0, 1, 1, 0, 2, 2, 0, 2, 0, 1, 1, 1, 0, 1, 0, 1, 1, 2, 2, 2, 2, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 1, 0, 0, 1, 0, 2, 1, 1, 1, 1, 0, 2, 2, 1, 0, 2, 0, 2, 2, 2, 1, 0, 2, 2, 2, 2, 0, 0, 1, 0, 1, 1, 2, 1, 2, 2, 1, 1, 0, 2, 0, 1, 0, 1, 1, 2, 0, 1, 1, 1, 1, 2, 0, 2, 2, 0, 0, 1, 1, 1, 1, 2, 1, 2, 2, 1, 2, 0, 1, 2, 2, 1, 1, 2, 1, 0, 1, 1, 1, 0, 0, 1, 2, 0, 2, 1, 0, 1, 2, 0, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0, 1, 0, 1, 0, 1, 2, 2, 2, 0, 1, 1, 1, 1, 1, 0, 2, 2, 2, 1, 0, 1, 0, 1, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 1, 0, 1, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 0, 0, 0, 0, 1, 2, 2, 1, 1, 0, 1, 2, 0, 0, 2, 1, 2, 1, 2, 0, 2, 1, 1, 2, 2, 2, 2, 2, 0, 1, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 2, 0, 1, 1, 2, 2, 0, 2, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 0, 2, 0, 2, 0, 2, 1, 0, 1, 0, 1, 1, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 0, 2, 2, 0, 2, 0, 0, 1, 1, 0, 2, 0, 1, 2, 2, 0, 1, 1, 2, 0, 0, 0, 2, 1, 0, 1, 0, 2, 1, 2, 0, 1, 2, 1, 1, 1, 0, 1, 2, 1, 2, 2, 1, 2, 0, 2, 1, 0, 1, 1, 1, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 1, 1, 1, 1, 0, 1, 2, 2, 0, 1, 2, 2, 1, 0, 0, 1, 2, 1, 0, 2, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 2, 0, 2, 0, 1, 2, 2, 0, 2, 1, 2, 2, 0, 0, 0, 0, 2, 1, 1, 2, 0, 2, 0, 1, 0, 1, 0, 2, 0, 0, 0, 2, 1, 2, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 0, 0, 1, 2, 1, 1, 2, 1, 1, 2, 2, 1, 0, 1, 1, 2, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 0, 0, 0, 2, 1, 1, 1, 2, 2, 1, 2, 0, 2, 1, 2, 0, 2, 2, 2, 1, 1, 2, 2, 0, 0, 2, 2, 2, 1, 2, 2, 1, 0, 2, 0, 2, 0, 2, 1, 2, 2, 1, 1, 1, 1, 1, 0, 0, 1, 0, 2, 2, 0, 0, 2, 1, 0, 1, 0, 2, 1, 0, 0, 0, 0, 1, 2, 1, 2, 0, 0, 1, 1, 2, 2, 2, 2, 0, 2, 1, 0, 0, 0, 2, 0, 1, 1, 0, 2, 1, 1, 1, 2, 1, 0, 0, 1, 0, 1, 0, 1, 2, 0, 2, 1, 0, 0, 1, 2, 1, 1, 0, 0, 0, 2, 2, 2, 1, 1, 2, 2, 0, 1, 0, 2, 2, 2, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 2, 1, 2, 0, 2, 0, 2, 1, 2, 2, 2, 0, 1, 0, 1, 2, 0, 2, 2, 1, 2, 0, 1, 2, 2, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 2, 1, 0, 0, 2, 0, 1, 0, 2, 2, 0, 1, 0, 0, 0, 1, 0, 0, 0, 2, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 2, 2, 0, 2, 2, 0, 0, 1, 2, 0, 1, 1, 1, 2, 0, 0, 0, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 0, 0, 1, 1, 2, 2, 2, 1, 1, 2, 0, 1, 0, 0, 1, 0, 1, 2, 0, 1, 2, 0, 1, 1, 1, 1, 2, 0, 1, 0, 1, 2, 2, 0, 0, 2, 0, 1, 2, 1, 2, 1, 0, 0, 1, 0, 0, 1, 2, 0, 1, 0, 0, 0, 0, 2, 0, 1, 0, 1, 2, 0, 1, 2, 0, 0, 0, 0, 1, 1, 2, 0, 0, 0, 2, 1, 1, 0, 0, 2, 2, 1, 0, 0, 1, 0, 1, 0, 1, 1, 0, 2, 0, 1, 1, 2, 2, 1, 1, 0, 2, 0, 0, 1, 0, 1, 2, 2, 1, 2, 2, 2, 2, 2, 1, 2, 0, 0, 0, 1, 1, 0, 0, 2, 1, 0, 1, 0, 2, 2, 0, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 2, 1, 0, 0, 0, 1, 1, 2, 1, 2, 0, 1, 2, 1, 0, 1, 2, 0, 1, 2, 1, 2, 1, 1, 1, 2, 2, 1, 0, 1, 0, 2, 1, 2, 0, 2, 0, 0, 0, 1, 2, 1, 0, 0, 0, 1, 2, 0, 2, 2, 2, 1, 0, 2, 2, 1, 1, 2, 0, 0, 1, 2, 2, 2, 1, 2, 1, 2, 1, 0, 2, 2, 0, 0, 0, 2, 0, 1, 0, 1, 0, 2, 2, 2, 2, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 2, 0, 1, 0, 0, 2, 0, 1, 1, 1, 1, 2, 0, 1, 1, 2, 0, 2, 1, 2, 1, 1, 0, 1, 1, 2, 1, 2, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 2, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 2, 2, 2, 0, 2, 0, 2, 1, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 1, 0, 2, 1, 2, 1, 1, 0, 2, 1, 2, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 2, 2, 1, 1, 0, 1, 0, 0, 0, 2, 2, 2, 0, 1, 2, 1, 2, 0, 2, 1, 0, 0, 1, 2, 2, 1, 0, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 1, 0, 2, 2, 2, 2, 0, 1, 0, 0, 0, 1, 1, 0, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 0, 2, 1, 0, 2, 2, 2, 1, 2, 0, 2, 2, 0, 2, 1, 2, 2, 0, 1, 0, 2, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 0, 2, 2, 2, 2, 1, 0, 2, 0, 2, 0, 1, 0, 1, 1, 2, 2, 1, 2, 1, 2, 0, 1, 2, 2, 2, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 0, 2, 1, 1, 1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 0, 0, 2, 2, 2, 2, 1, 2, 2, 1, 2, 0, 1, 0, 1, 1, 0, 0, 2, 2, 1, 2, 2, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 2, 1, 0, 1, 2, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 2, 2, 1, 2, 2, 2, 2, 1, 2, 2, 1, 1, 1, 0, 1, 2, 2, 0, 0, 2, 1, 0, 2, 1, 0, 2, 2, 0, 2, 1, 0, 2, 1, 0, 0, 0, 0, 2, 0, 1, 0, 2, 0, 1, 2, 2, 0, 1, 0, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 2, 0, 1, 1, 1, 0, 1, 2, 2, 1, 2, 2, 0, 2, 2, 0, 0, 0, 2, 2, 2, 1, 1, 2, 1, 2, 1, 2, 1, 1, 0, 0, 2, 0, 2, 0, 1, 1, 0, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 2, 1, 1, 1, 1, 2, 2, 0, 2, 2, 0, 1, 1, 2, 2, 0, 2, 0, 1, 0, 2, 0, 2, 2, 2, 1, 2, 2, 0, 1, 1, 2, 2, 1, 2, 2, 2, 2, 2, 2, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 2, 1, 2, 1, 2, 2, 1, 1, 0, 2, 1, 0, 0, 1, 2, 2, 1, 1, 1, 1, 0, 0, 0, 2, 0, 0, 2, 1, 1, 2, 2, 2, 1, 2, 1, 0, 1, 2, 1, 2, 2, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 2, 2, 1, 2, 1, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 1, 0, 2, 0, 2, 1, 1, 2, 0, 1, 1, 0, 2, 0, 1, 2, 2, 0, 0, 2, 0, 1, 1, 0, 2, 2, 1, 0, 0, 0, 2, 0, 1, 0, 0, 0, 2, 0, 2, 2, 1, 0, 1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 0, 1, 2, 0, 2, 1, 2, 1, 0, 2, 0, 1, 0, 0, 2, 1, 0, 0, 0, 2, 0, 0, 1, 2, 2, 0, 0, 1, 0, 1, 1, 2, 2, 2, 0, 2, 1, 0, 2, 0, 1, 2, 1, 0, 1, 1, 1, 2, 1, 0, 0, 0, 2, 1, 0, 2, 2, 0, 1, 0, 2, 0, 2, 2, 2, 0, 1, 2, 2, 2, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 2, 2, 2, 0, 1, 0, 0, 1, 2, 2, 2, 0, 1, 1, 1, 2, 0, 2, 1, 1, 2, 2, 1, 0, 1, 2, 0, 2, 1, 1, 0, 1, 1, 2, 0, 1, 0, 2, 1, 2, 0, 2, 2, 2, 2, 0, 2, 0, 2, 1, 0, 1, 0, 2, 1, 0, 0, 0, 2, 0, 0, 0, 1, 0, 2, 1, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 2, 1, 2, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 0, 1, 2, 2, 0, 0, 1, 1, 0, 2, 0, 2, 1, 2, 0, 0, 2, 1, 0, 0, 0, 2, 2, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 2, 1, 2, 1, 2, 0, 1, 1, 2, 0, 0, 2, 1, 0, 0, 2, 0, 1, 2, 1, 1, 2, 2, 2, 1, 2, 2, 0, 1, 2, 0, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 1, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 1, 1, 1, 2, 1, 2, 1, 2, 0, 1, 1, 1, 2, 0, 1, 2, 1, 1, 0, 2, 1, 0, 2, 0, 0, 0, 2, 1, 0, 1, 2, 0, 1, 0, 1, 1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 2, 2, 0, 1, 1, 2, 1, 2, 2, 1, 0, 1, 1, 2, 2, 0, 1, 0, 1, 0, 2, 1, 0, 2, 0, 2, 2, 1, 1, 1, 1, 2, 2, 2, 2, 0, 1, 1, 0, 2, 0, 1, 2, 1, 2, 0, 0, 2, 2, 1, 2, 2, 1, 1, 0, 0, 0, 0, 2, 1, 0, 2, 2, 0, 1, 2, 2, 2, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 0, 2, 0, 1, 2, 2, 1, 0, 2, 2, 0, 0, 0, 2, 0, 0, 0, 1, 2, 1, 0, 2, 2, 1, 2, 0, 0, 2, 2, 2, 1, 1, 1, 2, 0, 0, 2, 0, 0, 2, 1, 2, 2, 2, 1, 0, 0, 1, 1, 1, 0, 0, 2, 0, 1, 1, 0, 0, 0, 1, 0, 2, 0, 0, 1, 1, 2, 0, 2, 0, 2, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 0, 1, 2, 0, 2, 1, 0, 1, 1, 2, 0, 2, 2, 1, 1, 2, 2, 1, 1, 0, 2, 2, 2, 2, 1, 2, 0, 2, 0, 2, 2, 1, 1, 0, 1, 0, 0, 1, 2, 2, 1, 1, 1, 0, 1, 0, 1, 2, 0, 2, 1, 2, 2, 0, 2, 2, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 2, 0, 2, 2, 0, 1, 2, 2, 1, 1, 2, 0, 1, 1, 1, 2, 0, 2, 1, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1, 2, 0, 1, 1, 1, 0, 2, 1, 2, 0, 0, 2, 1, 0, 1, 0, 2, 2, 2, 1, 1, 1, 2, 2, 0, 2, 1, 0, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 2, 0, 0, 2, 0, 2, 2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 2, 0, 2, 2, 0, 2, 1, 0, 0, 1, 1, 1, 1, 2, 1, 2, 0, 2, 0, 0, 0, 2, 2, 0, 0, 2, 1, 2, 1, 0, 1, 0, 2, 0, 1, 2, 0, 1, 2, 0, 1, 1, 1, 1, 0, 2, 2, 1, 0, 2, 1, 0, 2, 0, 1, 0, 0, 2, 1, 0, 1, 0, 2, 2, 2, 1, 2, 2, 0, 2, 1, 0, 2, 0, 2, 1, 1, 2, 0, 1, 2, 0, 0, 1, 1, 1, 0, 2, 0, 0, 2, 2, 0, 0, 2, 2, 1, 0, 1, 2, 1, 1, 2, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 2, 2, 2, 1, 1, 2, 0, 1, 2, 1, 2, 1, 1, 1, 1, 1, 0, 2, 1, 1, 1, 0, 2, 0, 0, 0, 2, 1, 0, 1, 2, 0, 1, 1, 2, 1, 1, 2, 2, 0, 2, 0, 2, 0, 2, 0, 2, 2, 0, 0, 1, 0, 2, 0, 0, 0, 0, 0, 1, 0, 2, 1, 1, 0, 1, 2, 2, 1, 0, 2, 2, 1, 2, 0, 1, 2, 2, 2, 0, 1, 0, 0, 1, 0, 2, 1, 0, 1, 1, 1, 1, 2, 1, 2, 0, 1, 2, 2, 2, 2, 2, 0, 1, 0, 2, 2, 0, 0, 0, 0, 1, 2, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 1, 1, 1, 2, 0, 1, 1, 2, 0, 0, 1, 2, 0, 2, 0, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 2, 1, 1, 1, 2, 0, 1, 2, 2, 1, 1, 0, 2, 1, 2, 1, 1, 1, 0, 2, 1, 0, 0, 2, 2, 2, 1, 0, 1, 2, 1, 2, 1, 1, 0, 0, 0, 2, 1, 0, 1, 0, 0, 1, 2, 0, 0, 2, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 2, 2, 0, 0, 0, 1, 2, 0, 1, 1, 1, 1, 2, 1, 1, 0, 1, 1, 1, 0, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 1, 1, 0, 2, 1, 0, 2, 0, 1, 1, 0, 1, 2, 1, 1, 0, 2, 0, 0, 1, 0, 0, 1, 0, 0, 1, 2, 2, 0, 1, 2, 1, 2, 2, 2, 2, 0, 2, 1, 1, 0, 0, 0, 1, 1, 2, 2, 0, 0, 0, 1, 2, 0, 1, 2, 0, 1, 2, 1, 1, 2, 1, 2, 2, 0, 1, 2, 2, 1, 2, 0, 1, 1, 0, 0, 0, 0, 1, 2, 2, 2, 0, 1, 0, 1, 0, 2, 1, 1, 0, 2, 1, 2, 1, 2, 0, 2, 0, 0, 1, 2, 0, 1, 0, 0, 0, 1, 1, 1, 2, 1, 1, 1, 0, 1, 2, 0, 0, 0, 1, 0, 2, 0, 2, 0, 1, 1, 1, 0, 0, 0, 2, 0, 1, 2, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 2, 1, 1, 2, 2, 1, 1, 1, 0, 1, 0, 1, 1, 2, 1, 2, 2, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 0, 2, 1, 0, 0, 1, 1, 1, 0, 2, 2, 0, 0, 1, 1, 2, 2, 1, 2, 2, 2, 0, 0, 2, 2, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 0, 1, 2, 1, 2, 1, 2, 0, 0, 1, 0, 2, 0, 0, 2, 0, 1, 1, 2, 2, 0, 1, 0, 1, 1, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 0, 1, 1, 1, 0, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 2, 0, 0, 1, 0, 1, 0, 2, 2, 1, 0, 1, 2, 2, 0, 2, 2, 0, 0, 1, 1, 1, 2, 1, 2, 2, 2, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 0, 2, 1, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 1, 0, 0, 2, 0, 2, 2, 1, 0, 2, 0, 2, 2, 1, 1, 2, 2, 2, 1, 0, 0, 1, 2, 1, 0, 1, 0, 2, 2, 1, 2, 1, 2, 0, 1, 0, 1, 2, 2, 1, 0, 2, 0, 2, 1, 0, 1, 2, 0, 0, 2, 2, 1, 1, 0, 0, 1, 2, 0, 2, 2, 1, 2, 1, 0, 1, 1, 0, 0, 2, 2, 2, 2, 2, 0, 1, 0, 1, 1, 1, 0, 1, 2, 1, 0, 2, 2, 0, 1, 1, 2, 0, 2, 2, 0, 2, 1, 1, 2, 2, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 0, 2, 2, 1, 0, 0, 1, 1, 0, 1, 2, 1, 0, 0, 0, 2, 2, 0, 2, 0, 1, 2, 2, 1, 1, 1, 2, 0, 2, 0, 0, 1, 2, 0, 1, 1, 0, 1, 1, 1, 2, 0, 0, 0, 2, 1, 0, 2, 1, 2, 1, 0, 1, 1, 2, 2, 0, 1, 0, 2, 2, 2, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 0, 2, 1, 0, 0, 2, 0, 2, 0, 1, 1, 0, 0, 1, 2, 0, 2, 2, 2, 1, 1, 0, 0, 0, 0, 1, 0, 1, 2, 2, 2, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 2, 2, 0, 0, 2, 2, 0, 0, 0, 2, 2, 2, 2, 1, 1, 2, 2, 2, 0, 0, 2, 0, 1, 2, 2, 2, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 1, 1, 1, 1, 0, 2, 0, 2, 0, 0, 1, 2, 2, 1, 0, 2, 0, 1, 2, 2, 1, 2, 0, 0, 2, 0, 0, 1, 1, 1, 0, 2, 0, 2, 2, 2, 2, 1, 2, 2, 0, 1, 2, 2, 2, 0, 0, 0, 1, 0, 1, 2, 0, 0, 0, 0, 1, 2, 0, 0, 2, 2, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 2, 2, 0, 0, 1, 1, 0, 2, 2, 1, 1, 0, 2, 0, 2, 1, 0, 1, 2, 1, 1, 2, 1, 0, 0, 1, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 1, 1, 2, 0, 2, 1, 2, 0, 1, 0, 0, 0, 2, 0, 1, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 0, 2, 1, 0, 0, 2, 1, 1, 2, 0, 2, 0, 2, 0, 1, 2, 2, 0, 0, 0, 2, 2, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 2, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 0, 2, 1, 0, 1, 2, 0, 0, 2, 1, 1, 0, 0, 1, 0, 2, 1, 1, 1, 2, 1, 2, 2, 0, 2, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 2, 1, 1, 1, 1, 0, 0, 1, 1, 0, 2, 2, 0, 1, 1, 1, 1, 2, 1, 1, 0, 0, 2, 2, 0, 2, 1, 1, 1, 1, 2, 0, 2, 2, 2, 0, 1, 0, 1, 0, 0, 1, 0, 2, 0, 0, 0, 2, 1, 1, 2, 2, 0, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 0, 2, 2, 0, 0, 1, 2, 0, 1, 0, 0, 0, 0, 1, 2, 0, 1, 1, 0, 0, 2, 0, 0, 2, 0, 1, 1, 2, 2, 1, 0, 1, 0, 0, 2, 2, 2, 0, 0, 1, 0, 2, 0, 1, 2, 1, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 2, 1, 1, 2, 0, 1, 1, 1, 0, 0, 0, 1, 0, 2, 0, 2, 0, 2, 0, 0, 1, 0, 0, 2, 0, 2, 2, 2, 2, 2, 1, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 0, 1, 0, 2, 2, 2, 1, 1, 1, 0, 1, 0, 1, 1, 0, 2, 0, 0, 1, 2, 2, 1, 2, 0, 0, 0, 2, 2, 0, 1, 1, 0, 2, 0, 1, 2, 1, 0, 0, 0, 1, 1, 1, 2, 0, 1, 1, 1, 2, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 1, 2, 0, 2, 0, 0, 0, 2, 0, 2, 2, 2, 0, 1, 1, 2, 0, 0, 1, 1, 0, 1, 1, 0, 2, 0, 1, 1, 2, 2, 1, 0, 2, 1, 2, 0, 0, 0, 2, 2, 2, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 2, 0, 2, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 2, 0, 0, 2, 0, 2, 2, 2, 2, 1, 1, 1, 2, 0, 2, 1, 1, 0, 0, 1, 2, 1, 2, 2, 2, 2, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 2, 0, 2, 2, 2, 0, 1, 0, 1, 2, 2, 0, 0, 2, 0, 0, 1, 0, 0, 1, 2, 1, 0, 1, 2, 2, 1, 2, 2, 1, 1, 0, 2, 1, 0, 1, 1, 2, 1, 2, 1, 0, 2, 0, 1, 1, 0, 0, 1, 0, 2, 1, 1, 1, 2, 1, 0, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, 0, 0, 2, 2, 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 2, 2, 1, 0, 1, 0, 1, 1, 2, 2, 0, 1, 1, 2, 2, 1, 0, 2, 1, 1, 1, 2, 1, 2, 0, 0, 2, 0, 2, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 2, 0, 2, 1, 0, 2, 1, 1, 0, 1, 0, 0, 2, 2, 1, 1, 2, 0, 0, 0, 1, 1, 1, 1, 1, 0, 2, 2, 2, 2, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 2, 2, 1, 0, 1, 1, 0, 2, 0, 2, 2, 1, 1, 1, 1, 0, 0, 2, 2, 1, 2, 0, 0, 2, 1, 0, 2, 1, 2, 1, 2, 1, 0, 2, 2, 2, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 2, 2, 2, 1, 0, 2, 0, 2, 2, 2, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 2, 1, 2, 0, 0, 0, 1, 2, 0, 1, 2, 0, 1, 0, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 1, 0, 2, 1, 0, 1, 2, 1, 1, 1, 2, 1, 0, 0, 2, 1, 1, 1, 0, 2, 1, 2, 2, 2, 2, 2, 1, 1, 2, 1, 2, 1, 0, 1, 1, 0, 1, 2, 0, 1, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 2, 0, 1, 2, 1, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 2, 0, 1, 1, 2, 0, 0, 2, 2, 0, 0, 2, 0, 2, 2, 1, 0, 1, 0, 2, 0, 0, 0, 0, 2, 0, 1, 2, 1, 0, 1, 2, 0, 2, 0, 2, 1, 0, 0, 0, 1, 2, 0, 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 2, 1, 2, 1, 1, 1, 2, 0, 2, 1, 0, 0, 0, 2, 2, 1, 0, 2, 0, 2, 0, 1, 0, 1, 0, 2, 0, 1, 1, 1, 1, 0, 2, 0, 1, 2, 1, 2, 0, 2, 1, 2, 1, 2, 2, 2, 2, 1, 2, 0, 0, 2, 0, 2, 2, 2, 0, 1, 2, 1, 0, 2, 0, 2, 2, 2, 0, 1, 1, 0, 2, 2, 2, 2, 0, 1, 2, 1, 2, 2, 1, 2, 0, 1, 2, 0, 0, 0, 2, 2, 1, 1, 0, 0, 2, 1, 1, 1, 0, 1, 2, 0, 1, 2, 0, 0, 1, 2, 0, 2, 0, 0, 2, 0, 2, 0, 1, 1, 1, 1, 1, 0, 2, 1, 1, 2, 2, 0, 0, 1, 2, 1, 2, 2, 1, 0, 2, 2, 2, 2, 2, 0, 1, 0, 0, 1, 1, 2, 0, 1, 2, 1, 2, 0, 0, 0, 2, 2, 2, 1, 2, 1, 1, 2, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 2, 0, 1, 1, 2, 2, 0, 0, 1, 2, 0, 2, 2, 1, 2, 0, 0, 1, 0, 1, 2, 2, 0, 1, 2, 1, 2, 1, 2, 1, 2, 0, 1, 1, 1, 1, 1, 2, 0, 2, 2, 1, 2, 2, 1, 1, 0, 0, 1, 0, 2, 0, 1, 1, 1, 1, 2, 2, 0, 0, 1, 0, 1, 0, 2, 2, 2, 1, 0, 2, 2, 2, 0, 0, 2, 2, 1, 2, 2, 0, 2, 0, 2, 1, 0, 2, 0, 0, 1, 1, 1, 0, 2, 0, 0, 0, 1, 0, 2, 1, 2, 1, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 1, 0, 1, 2, 1, 2, 1, 0, 2, 2, 2, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0, 0, 2, 0, 1, 2, 2, 0, 1, 2, 2, 2, 2, 0, 2, 1, 1, 2, 2, 2, 2, 1, 0, 1, 0, 2, 2, 0, 1, 0, 2, 2, 0, 0, 0, 0, 2, 0, 2, 1, 0, 1, 1, 2, 2, 1, 1, 2, 1, 2, 2, 0, 1, 1, 2, 0, 0, 0, 0, 2, 0, 0, 1, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 0, 2, 1, 0, 2, 2, 1, 0, 0, 2, 1, 2, 2, 2, 1, 0, 2, 1, 0, 1, 2, 2, 0, 1, 2, 2, 0, 0, 2, 1, 2, 0, 1, 0, 1, 2, 1, 1, 0, 1, 1, 0, 0, 1, 2, 0, 1, 2, 2, 0, 1, 0, 0, 2, 2, 2, 2, 2, 0, 2, 0, 1, 1, 2, 2, 0, 0, 2, 0, 0, 2, 0, 1, 1, 0, 1, 0, 1, 2, 1, 1, 2, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 1, 0, 1, 2, 2, 0, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 0, 0, 0, 0, 1, 1, 0, 0, 2, 2, 1, 2, 0, 2, 2, 0, 0, 1, 1, 0, 2, 1, 1, 1, 0, 2, 0, 0, 0, 0, 0, 1, 2, 1, 0, 0, 0, 0, 2, 2, 1, 0, 2, 2, 0, 1, 0, 0, 1, 2, 0, 0, 0, 1, 0, 2, 0, 2, 0, 0, 2, 2, 1, 1, 0, 0, 2, 0, 0, 2, 1, 0, 2, 1, 0, 1, 2, 2, 1, 0, 2, 0, 0, 0, 0, 2, 2, 2, 2, 0, 1, 1, 0, 1, 2, 0, 0, 0, 2, 0, 1, 1, 1, 1, 0, 0, 1, 2, 1, 1, 2, 2, 1, 0, 1, 1, 1, 1, 0, 2, 0, 1, 2, 2, 0, 2, 0, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 1, 1, 2, 1, 1, 2, 0, 0, 2, 0, 0, 1, 0, 0, 2, 2, 0, 2, 1, 2, 2, 2, 2, 1, 0, 1, 2, 2, 1, 1, 1, 1, 1, 0, 1, 2, 0, 0, 2, 0, 0, 2, 2, 1, 1, 1, 0, 0, 0, 2, 0, 2, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 0, 1, 2, 2, 2, 1, 0, 2, 1, 1, 2, 0, 1, 2, 2, 1, 0, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 0, 0, 2, 2, 1, 1, 0, 2, 0, 0, 2, 2, 0, 1, 0, 1, 2, 0, 1, 0, 1, 2, 0, 0, 1, 1, 1, 0, 1, 1, 0, 2, 2, 2, 0, 1, 0, 1, 1, 0, 2, 1, 0, 1, 0, 0, 2, 0, 0, 1, 2, 0, 2, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 2, 0, 0, 2, 1, 2, 1, 2, 2, 1, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 0, 2, 2, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 2, 1, 0, 1, 0, 1, 0, 2, 2, 2, 2, 2, 1, 2, 0, 1, 2, 2, 1, 0, 1, 2, 0, 1, 1, 0, 1, 2, 0, 1, 2, 0, 0, 2, 0, 0, 2, 2, 0, 0, 1, 2, 2, 0, 0, 1, 2, 2, 1, 1, 1, 2, 0, 2, 0, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2, 2, 2, 2, 2, 0, 0, 1, 2, 1, 2, 0, 0, 2, 0, 0, 2, 1, 1, 1, 2, 0, 1, 2, 0, 1, 1, 0, 1, 1, 2, 0, 2, 2, 0, 1, 2, 2, 2, 0, 1, 2, 0, 2, 2, 0, 0, 1, 0, 0, 0, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 1, 0, 2, 0, 1, 1, 1, 0, 2, 0, 0, 1, 0, 2, 2, 1, 0, 0, 1, 1, 0, 1, 2, 1, 1, 2, 2, 2, 0, 1, 0, 2, 2, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 1, 1, 2, 0, 1, 2, 1, 2, 1, 2, 2, 1, 2, 2, 0, 1, 1, 2, 0, 1, 0, 0, 2, 0, 0, 1, 0, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 2, 1, 1, 0, 0, 1, 2, 2, 0, 1, 2, 0, 1, 0, 2, 2, 1, 1, 1, 0, 0, 2, 2, 2, 2, 0, 1, 0, 2, 2, 1, 2, 1, 1, 1, 0, 2, 2, 1, 1, 2, 2, 0, 2, 1, 2, 2, 2, 1, 1, 2, 1, 0, 0, 1, 1, 0, 2, 1, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 2, 1, 0, 2, 1, 0, 0, 2, 0, 2, 2, 1, 2, 0, 2, 1, 2, 2, 1, 2, 0, 0, 2, 2, 0, 0, 1, 0, 1, 2, 1, 1, 0, 2, 1, 0, 1, 2, 2, 2, 1, 1, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 2, 0, 0, 0, 0, 0, 2, 0, 2, 1, 0, 0, 0, 1, 2, 0, 0, 2, 1, 2, 0, 2, 1, 2, 0, 1, 1, 2, 0, 0, 1, 1, 0, 0, 2, 0, 1, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 1, 2, 0, 2, 1, 2, 2, 1, 1, 1, 0, 1, 2, 2, 1, 2, 0, 0, 0, 2, 1, 1, 1, 0, 2, 0, 1, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 2, 0, 1, 0, 0, 2, 2, 1, 0, 0, 0, 2, 2, 0, 2, 1, 1, 0, 1, 1, 2, 1, 2, 2, 0, 0, 1, 2, 2, 2, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 2, 2, 2, 2, 1, 2, 1, 2, 0, 0, 1, 1, 2, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 0, 0, 1, 2, 0, 1, 1, 0, 2, 2, 2, 0, 1, 2, 2, 1, 0, 1, 2, 1, 0, 2, 2, 2, 2, 0, 1, 0, 2, 0, 0, 0, 1, 1, 2, 2, 2, 0, 2, 2, 2, 0, 2, 1, 0, 2, 1, 2, 1, 2, 2, 1, 0, 1, 1, 2, 2, 0, 1, 2, 1, 2, 1, 2, 2, 0, 2, 0, 1, 1, 0, 2, 2, 2, 2, 1, 0, 1, 0, 0, 1, 1, 2, 1, 2, 1, 1, 2, 1, 1, 0, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 2, 2, 1, 2, 1, 2, 0, 1, 2, 1, 0, 2, 0, 2, 0, 0, 0, 2, 0, 0, 0, 1, 2, 2, 2, 2, 1, 0, 2, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 2, 2, 0, 2, 1, 2, 1, 2, 2, 1, 0, 1, 0, 2, 2, 0, 2, 2, 1, 1, 0, 0, 1, 1, 0, 2, 2, 2, 0, 1, 0, 1, 0, 0, 1, 0, 2, 1, 2, 2, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 2, 2, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 2, 0, 2, 2, 0, 2, 2, 2, 0, 1, 2, 0, 2, 2, 2, 0, 1, 0, 0, 2, 2, 1, 1, 2, 1, 2, 2, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 2, 0, 0, 2, 1, 2, 0, 2, 2, 2, 2, 2, 2, 1, 2, 0, 2, 2, 2, 2, 1, 1, 0, 0, 1, 2, 0, 2, 0, 2, 0, 2, 2, 0, 1, 2, 1, 2, 0, 2, 2, 1, 2, 1, 0, 0, 1, 2, 1, 0, 2, 2, 1, 1, 1, 2, 1, 1, 2, 1, 0, 2, 2, 0, 0, 0, 2, 1, 2, 2, 0, 1, 0, 0, 2, 1, 1, 1, 0, 1, 1, 2, 2, 1, 1, 0, 0, 1, 1, 0, 2, 0, 2, 1, 0, 0, 1, 2, 2, 1, 0, 2, 2, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 2, 1, 0, 2, 2, 0, 2, 1, 1, 0, 2, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 2, 0, 2, 2, 2, 1, 1, 1, 0, 2, 1, 2, 1, 0, 1, 1, 0, 0, 0, 2, 1, 2, 2, 1, 1, 2, 0, 1, 0, 2, 0, 0, 2, 2, 1, 0, 2, 1, 2, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 2, 1, 1, 0, 2, 1, 2, 2, 1, 2, 2, 1, 0, 1, 2, 2, 0, 1, 0, 1, 2, 2, 1, 1, 1, 0, 0, 0, 1, 2, 2, 1, 2, 2, 0, 2, 2, 1, 0, 0, 2, 0, 1, 1, 2, 1, 0, 0, 0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 2, 0, 2, 1, 2, 0, 1, 1, 1, 0, 0, 2, 2, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 1, 1, 0, 1, 0, 0, 2, 0, 1, 1, 2, 2, 2, 1, 1, 0, 1, 2, 2, 2, 0, 2, 2, 0, 1, 1, 1, 1, 1, 2, 0, 0, 2, 2, 2, 0, 0, 0, 1, 1, 2, 0, 2, 1, 0, 1, 1, 2, 1, 1, 2, 2, 2, 2, 2, 2, 0, 0, 2, 1, 2, 2, 2, 1, 2, 0, 0, 2, 1, 1, 0, 1, 0, 2, 0, 1, 2, 2, 0, 0, 0, 2, 0, 1, 0, 1, 2, 1, 0, 1, 2, 2, 1, 0, 0, 1, 1, 0, 2, 2, 0, 1, 2, 1, 1, 0, 1, 1, 2, 1, 2, 2, 1, 0, 2, 0, 2, 1, 0, 0, 2, 2, 1, 1, 0, 0, 1, 1, 0, 0, 2, 0, 0, 2, 0, 2, 0, 2, 1, 1, 2, 1, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 0, 0, 2, 1, 1, 2, 0, 2, 1, 1, 1, 2, 1, 2, 1, 2, 0, 1, 2, 2, 2, 1, 2, 2, 1, 1, 2, 0, 2, 1, 0, 1, 2, 1, 1, 0, 2, 1, 2, 0, 1, 1, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 1, 2, 0, 0, 1, 0, 0, 2, 2, 0, 2, 1, 0, 0, 1, 1, 2, 2, 2, 0, 1, 1, 0, 1, 0, 2, 0, 0, 0, 2, 0, 1, 0, 0, 0, 1, 0, 2, 2, 0, 1, 1, 1, 1, 0, 0, 2, 1, 2, 0, 0, 0, 1, 1, 2, 2, 0, 0, 0, 1, 2, 1, 0, 1, 1, 2, 1, 1, 0, 0, 2, 2, 2, 0, 1, 0, 0, 2, 1, 0, 1, 0, 0, 0, 2, 1, 2, 1, 0, 2, 0, 0, 1, 2, 2, 1, 2, 1, 2, 0, 0, 1, 1, 0, 1, 1, 1, 1, 2, 2, 1, 2, 0, 2, 1, 0, 0, 0, 2, 2, 2, 0, 1, 2, 1, 1, 0, 1, 2, 1, 0, 0, 0, 2, 0, 0, 1, 2, 2, 0, 0, 1, 0, 0, 1, 0, 2, 1, 0, 2, 2, 2, 0, 2, 1, 1, 2, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 2, 0, 2, 2, 2, 2, 2, 0, 2, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 2, 1, 2, 2, 2, 1, 0, 1, 0, 2, 2, 0, 0, 0, 1, 2, 0, 2, 2, 1, 1, 1, 0, 2, 2, 0, 2, 2, 0, 1, 2, 2, 0, 1, 1, 1, 1, 2, 2, 0, 2, 2, 2, 2, 1, 1, 0, 2, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 2, 1, 2, 0, 1, 0, 2, 1, 2, 2, 2, 2, 0, 2, 1, 0, 2, 0, 2, 0, 2, 2, 0, 0, 2, 0, 2, 1, 2, 2, 0, 0, 1, 2, 0, 1, 1, 2, 0, 0, 2, 0, 1, 0, 0, 2, 0, 1, 2, 2, 1, 0, 2, 2, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 0, 2, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 2, 1, 1, 0, 1, 0, 1, 1, 2, 0, 2, 1, 0, 0, 0, 2, 0, 2, 0, 0, 2, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 2, 0, 1, 2, 1, 0, 1, 0, 2, 2, 2, 2, 0, 2, 0, 0, 0, 1, 1, 0, 1, 1, 0, 2, 0, 2, 2, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 2, 1, 0, 1, 2, 1, 2, 2, 0, 1, 0, 1, 2, 0, 1, 1, 2, 0, 0, 2, 0, 2, 0, 0, 1, 2, 2, 2, 0, 0, 2, 0, 1, 1, 2, 2, 0, 0, 2, 1, 1, 0, 1, 1, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 1, 2, 1, 2, 0, 2, 1, 1, 1, 0, 2, 0, 0, 2, 2, 2, 2, 2, 1, 2, 0, 2, 1, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 0, 1, 0, 0, 1, 1, 0, 1, 1, 2, 0, 2, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 2, 0, 0, 0, 2, 2, 0, 0, 0, 0, 2, 1, 2, 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 2, 1, 2, 1, 0, 1, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1, 0, 0, 1, 1, 2, 1, 2, 1, 0, 1, 1, 1, 1, 2, 1, 1, 2, 0, 1, 2, 1, 0, 2, 2, 2, 0, 0, 2, 2, 0, 0, 2, 1, 0, 0, 2, 2, 1, 2, 0, 0, 0, 0, 2, 1, 0, 1, 1, 0, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 1, 2, 0, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 1, 1, 2, 1, 1, 2, 1, 2, 1, 2, 2, 0, 2, 2, 1, 2, 1, 2, 2, 2, 0, 0, 0, 1, 1, 0, 2, 1, 0, 1, 1, 2, 1, 0, 2, 1, 0, 1, 0, 0, 2, 2, 1, 2, 1, 0, 1, 1, 0, 1, 1, 2, 2, 0, 0, 2, 0, 0, 0, 2, 0, 2, 1, 1, 2, 2, 1, 0, 0, 2, 1, 0, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 0, 1, 1, 1, 2, 0, 2, 0, 0, 1, 2, 0, 2, 2, 0, 2, 2, 0, 0, 2, 1, 2, 1, 1, 0, 2, 2, 2, 1, 2, 2, 0, 0, 0, 0, 2, 0, 1, 1, 2, 1, 1, 1, 2, 2, 1, 0, 1, 1, 1, 1, 0, 1, 1, 2, 1, 0, 2, 0, 2, 0, 1, 2, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 1, 0, 1, 2, 1, 2, 0, 2, 1, 2, 1, 0, 1, 1, 1, 0, 1, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 2, 2, 0, 0, 2, 2, 0, 2, 1, 2, 0, 1, 0, 2, 1, 1, 2, 1, 2, 0, 2, 1, 2, 1, 1, 2, 2, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 2, 1, 2, 0, 1, 1, 2, 0, 2, 2, 0, 2, 0, 1, 1, 1, 0, 0, 2, 1, 1, 2, 0, 1, 2, 0, 0, 2, 0, 0, 2, 0, 1, 2, 2, 2, 2, 2, 0, 1, 0, 2, 2, 1, 1, 2, 2, 2, 1, 0, 2, 2, 2, 2, 1, 0, 0, 2, 2, 1, 1, 1, 1, 1, 1, 2, 2, 0, 2, 2, 1, 1, 2, 0, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 2, 1, 1, 1, 2, 0, 0, 0, 1, 2, 0, 1, 0, 1, 0, 1, 1, 2, 2, 1, 0, 0, 2, 0, 1, 2, 0, 0, 1, 2, 1, 1, 2, 2, 1, 0, 0, 0, 0, 2, 2, 0, 2, 0, 0, 1, 0, 2, 0, 2, 1, 1, 1, 0, 2, 2, 0, 0, 0, 1, 0, 1, 0, 1, 0, 2, 1, 2, 0, 1, 2, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 1, 0, 2, 1, 1, 1, 2, 0, 1, 1, 2, 1, 0, 1, 2, 0, 2, 1, 2, 0, 0, 1, 1, 0, 1, 1, 2, 0, 2, 2, 0, 2, 1, 0, 0, 2, 2, 0, 2, 1, 2, 1, 1, 1, 1, 2, 2, 0, 1, 2, 2, 0, 2, 1, 1, 1, 2, 0, 0, 1, 0, 0, 0, 2, 1, 1, 0, 0, 0, 1, 0, 0, 2, 2, 0, 1, 1, 0, 2, 1, 2, 2, 0, 0, 2, 2, 0, 0, 2, 2, 0, 0, 1, 1, 1, 0, 0, 0, 1, 2, 0, 0, 0, 1, 0, 1, 2, 2, 2, 1, 1, 1, 1, 0, 2, 2, 1, 1, 1, 0, 2, 0, 2, 0, 2, 1, 2, 2, 1, 2, 1, 0, 1, 1, 2, 1, 2, 2, 1, 0, 1, 1, 0, 1, 2, 1, 0, 0, 0, 0, 0, 2, 1, 2, 0, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 1, 2, 0, 2, 2, 1, 1, 1, 0, 0, 2, 0, 1, 1, 1, 1, 2, 0, 1, 1, 0, 0, 2, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 2, 1, 2, 1, 1, 2, 2, 2, 1, 1, 0, 2, 2, 0, 2, 2, 2, 0, 2, 0, 1, 1, 0, 1, 2, 1, 2, 1, 1, 2, 0, 0, 1, 2, 1, 1, 0, 1, 2, 1, 2, 2, 1, 2, 1, 1, 0, 2, 2, 1, 0, 2, 0, 1, 2, 1, 1, 2, 0, 1, 2, 2, 1, 2, 0, 1, 2, 1, 1, 2, 0, 2, 1, 1, 0, 2, 2, 2, 0, 2, 0, 1, 1, 0, 0, 1, 0, 1, 2, 2, 1, 1, 0, 1, 1, 1, 1, 1, 2, 2, 0, 2, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 1, 2, 1, 0, 0, 0, 1, 0, 0, 1, 2, 1, 1, 0, 2, 0, 0, 0, 0, 0, 1, 0, 1, 2, 1, 2, 2, 2, 0, 2, 1, 1, 2, 2, 1, 0, 1, 0, 0, 1, 1, 2, 1, 2, 0, 1, 2, 2, 2, 0, 2, 2, 0, 0, 2, 1, 2, 1, 2, 1, 0, 2, 2, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 2, 2, 2, 0, 1, 2, 0, 1, 1, 2, 1, 0, 1, 2, 0, 2, 2, 2, 1, 2, 2, 1, 1, 2, 0, 2, 1, 1, 0, 2, 1, 1, 1, 2, 1, 2, 1, 1, 2, 1, 2, 2, 1, 0, 0, 2, 0, 0, 1, 0, 2, 0, 2, 0, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 1, 0, 2, 2, 1, 2, 1, 1, 0, 2, 2, 2, 2, 0, 2, 1, 0, 2, 1, 2, 2, 0, 1, 1, 1, 0, 1, 2, 1, 2, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 1, 0, 2, 1, 0, 0, 2, 1, 0, 1, 0, 0, 1, 1, 1, 1, 2, 0, 1, 0, 1, 1, 2, 2, 2, 0, 0, 2, 1, 1, 1, 2, 2, 0, 0, 0, 0, 2, 0, 1, 1, 1, 2, 0, 0, 1, 1, 1, 2, 0, 0, 1, 0, 2, 0, 0, 2, 2, 1, 2, 1, 1, 2, 2, 2, 2, 1, 2, 2, 1, 2, 1, 1, 0, 1, 1, 1, 1, 2, 1, 1, 2, 2, 2, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 1, 1, 0, 1, 1, 2, 2, 0, 1, 2, 2, 1, 2, 1, 0, 1, 2, 1, 0, 2, 1, 1, 1, 2, 2, 1, 1, 0, 0, 0, 1, 1, 2, 1, 0, 2, 1, 0, 1, 2, 1, 1, 1, 1, 2, 2, 0, 2, 1, 1, 1, 1, 2, 0, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 0, 0, 0, 1, 2, 0, 1, 2, 0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 0, 0, 1, 0, 0, 2, 0, 1, 0, 2, 0, 1, 2, 1, 1, 1, 1, 2, 0, 2, 0, 2, 0, 0, 1, 2, 2, 1, 1, 1, 1, 1, 2, 1, 2, 0, 0, 2, 1, 2, 1, 1, 1, 2, 2, 2, 1, 2, 0, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 0, 0, 0, 2, 0, 2, 2, 0, 1, 0, 1, 2, 1, 1, 0, 0, 1, 1, 2, 2, 0, 1, 1, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 2, 1, 0, 1, 0, 1, 1, 0, 2, 2, 2, 1, 2, 0, 2, 2, 2, 2, 2, 1, 0, 1, 1, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2, 0, 2, 2, 1, 1, 1, 1, 2, 2, 0, 0, 0, 2, 0, 2, 1, 0, 0, 1, 0, 1, 2, 1, 0, 1, 0, 2, 0, 2, 0, 0, 2, 2, 0, 0, 1, 2, 1, 0, 1, 0, 2, 1, 0, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 2, 0, 2, 2, 2, 1, 1, 2, 2, 1, 0, 0, 1, 1, 2, 1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 2, 0, 2, 2, 1, 2, 2, 2, 0, 0, 2, 1, 1, 2, 1, 2, 2, 0, 0, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 2, 1, 0, 0, 1, 0, 2, 2, 1, 0, 2, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 0, 2, 1, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 2, 0, 1, 0, 2, 1, 2, 0, 2, 1, 0, 2, 0, 1, 1, 0, 1, 1, 2, 2, 0, 0, 2, 1, 0, 2, 0, 2, 1, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 0, 2, 1, 2, 2, 1, 0, 2, 0, 0, 1, 0, 2, 0, 1, 1, 1, 2, 2, 2, 2, 2, 0, 1, 1, 1, 1, 2, 1, 0, 0, 2, 2, 1, 2, 2, 2, 0, 1, 2, 1, 1, 1, 2, 2, 0, 1, 1, 2, 2, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 0, 1, 2, 1, 1, 0, 2, 1, 2, 1, 0, 0, 0, 0, 1, 2, 1, 0, 0, 0, 0, 1, 2, 0, 0, 1, 2, 2, 1, 2, 2, 2, 2, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 0, 0, 0, 2, 1, 0, 2, 1, 0, 2, 0, 2, 2, 0, 0, 0, 0, 2, 1, 1, 1, 0, 1, 2, 2, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 2, 1, 2, 1, 1, 2, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 1, 2, 1, 2, 1, 0, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 0, 0, 2, 0, 1, 1, 0, 1, 1, 1, 2, 1, 0, 0, 1, 2, 0, 1, 1, 0, 2, 1, 1, 2, 1, 0, 1, 2, 0, 1, 1, 0, 1, 1, 0, 2, 2, 1, 0, 1, 1, 0, 0, 0, 2, 2, 2, 1, 0, 2, 0, 1, 0, 2, 2, 1, 2, 0, 0, 1, 2, 1, 0, 1, 0, 1, 2, 2, 2, 2, 1, 2, 1, 0, 2, 1, 2, 2, 2, 2, 2, 0, 2, 2, 0, 1, 1, 2, 2, 2, 0, 2, 2, 2, 0, 0, 2, 1, 1, 1, 0, 0, 1, 0, 0, 2, 1, 1, 1, 0, 1, 0, 0, 0, 2, 0, 0, 1, 1, 1, 2, 0, 1, 1, 0, 1, 2, 0, 1, 2, 0, 2, 2, 2, 1, 0, 2, 2, 2, 1, 0, 0, 2, 2, 2, 0, 2, 0, 2, 1, 1, 0, 1, 0, 0, 1, 2, 1, 1, 2, 1, 1, 0, 2, 0, 0, 1, 1, 0, 2, 1, 2, 1, 1, 1, 2, 0, 2, 1, 0, 1, 0, 1, 0, 1, 0, 1, 2, 0, 1, 0, 0, 2, 2, 1, 2, 0, 0, 2, 1, 2, 1, 2, 2, 0, 1, 1, 0, 1, 1, 2, 0, 0, 1, 2, 0, 0, 0, 2, 0, 2, 2, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 1, 2, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 2, 2, 1, 0, 1, 1, 1, 1, 2, 0, 0, 1, 0, 2, 1, 0, 2, 1, 1, 0, 1, 1, 0, 1, 2, 2, 0, 0, 1, 1, 2, 0, 2, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 0, 0, 1, 0, 2, 0, 2, 0, 0, 1, 0, 0, 1, 2, 0, 1, 1, 1, 2, 0, 0, 0, 0, 2, 2, 2, 0, 1, 0, 1, 1, 0, 0, 2, 2, 0, 2, 0, 1, 2, 1, 1, 1, 0, 1, 0, 2, 0, 1, 2, 0, 2, 2, 0, 1, 1, 1, 2, 1, 1, 0, 1, 2, 2, 0, 1, 0, 1, 0, 1, 2, 1, 1, 0, 0, 2, 1, 2, 2, 2, 0, 1, 2, 0, 2, 1, 1, 0, 2, 1, 2, 1, 2, 0, 2, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 1, 2, 0, 1, 0, 0, 0, 1, 0, 2, 2, 1, 0, 1, 1, 0, 2, 2, 1, 0, 2, 2, 1, 1, 0, 0, 0, 1, 2, 2, 1, 2, 0, 2, 0, 1, 0, 2, 2, 0, 0, 2, 0, 1, 2, 2, 1, 1, 2, 2, 1, 1, 1, 2, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 2, 0, 0, 1, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 1, 2, 1, 2, 0, 1, 2, 2, 0, 2, 1, 1, 0, 0, 0, 2, 0, 2, 0, 1, 2, 0, 1, 2, 1, 0, 1, 1, 0, 0, 0, 1, 2, 1, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 2, 1, 0, 2, 1, 1, 0, 1, 1, 2, 1, 1, 1, 0, 2, 2, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 2, 0, 0, 2, 2, 1, 0, 2, 2, 2, 2, 2, 0, 1, 0, 1, 1, 2, 1, 1, 2, 1, 1, 1, 2, 1, 1, 2, 1, 0, 0, 2, 2, 0, 0, 1, 1, 2, 1, 1, 0, 0, 0, 2, 0, 2, 2, 1, 0, 2, 0, 0, 1, 2, 1, 2, 2, 1, 2, 0, 0, 2, 0, 0, 0, 2, 2, 1, 0, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0, 2, 0, 0, 0, 0, 1, 0, 0, 1, 2, 1, 2, 2, 2, 0, 2, 2, 0, 1, 2, 0, 1, 0, 2, 0, 2, 1, 1, 1, 0, 2, 1, 2, 1, 2, 0, 1, 0, 2, 2, 2, 0, 1, 1, 1, 2, 2, 2, 0, 1, 0, 2, 2, 2, 2, 0, 0, 0, 2, 1, 2, 2, 1, 2, 0, 1, 1, 1, 0, 0, 1, 1, 1, 2, 1, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 1, 0, 2, 1, 2, 1, 1, 2, 1, 1, 2, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 0, 0, 2, 1, 2, 0, 2, 2, 1, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 0, 0, 2, 0, 2, 2, 1, 2, 0, 2, 1, 0, 1, 2, 2, 2, 0, 2, 2, 2, 2, 1, 2, 1, 2, 2, 0, 0, 2, 2, 0, 1, 2, 0, 2, 1, 0, 2, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 0, 1, 1, 0, 0, 1, 0, 2, 2, 0, 1, 2, 1, 0, 2, 2, 0, 1, 0, 2, 2, 1, 1, 0, 2, 1, 1, 2, 2, 1, 2, 2, 1, 0, 1, 2, 0, 1, 1, 1, 0, 2, 0, 0, 2, 2, 0, 2, 1, 0, 2, 0, 1, 0, 0, 1, 0, 1, 0, 2, 2, 1, 1, 0, 1, 0, 1, 0, 1, 2, 2, 2, 2, 1, 2, 1, 0, 0, 2, 1, 0, 0, 0, 1, 2, 2, 0, 1, 1, 2, 2, 0, 1, 0, 0, 1, 0, 1, 1, 1, 2, 0, 2, 0, 0, 1, 1, 1, 0, 1, 1, 2, 2, 2, 1, 0, 0, 2, 2, 0, 2, 2, 1, 2, 0, 0, 2, 2, 0, 0, 2, 2, 0, 0, 1, 0, 2, 1, 0, 1, 1, 0, 2, 0, 0, 0, 0, 2, 1, 1, 2, 0, 0, 2, 1, 0, 0, 1, 2, 1, 1, 0, 0, 0, 2, 2, 2, 1, 2, 1, 0, 1, 1, 1, 0, 1, 0, 0, 2, 0, 0, 1, 0, 2, 2, 2, 2, 2, 1, 1, 0, 0, 0, 1, 1, 0, 2, 1, 1, 1, 2, 1, 0, 0, 1, 0, 2, 2, 2, 0, 1, 1, 0, 1, 1, 2, 0, 1, 2, 1, 0, 2, 1, 0, 0, 2, 1, 0, 2, 1, 2, 2, 1, 2, 0, 2, 2, 0, 0, 1, 0, 2, 2, 1, 2, 1, 2, 2, 1, 1, 0, 0, 2, 2, 1, 1, 2, 2, 1, 2, 0, 0, 0, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 2, 2, 2, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 2, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 2, 0, 0, 1, 1, 2, 0, 2, 1, 0, 0, 2, 0, 1, 2, 0, 0, 0, 1, 1, 0, 2, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 0, 2, 0, 2, 1, 0, 1, 0, 1, 0, 1, 2, 1, 1, 0, 0, 1, 2, 1, 0, 2, 0, 1, 1, 1, 2, 2, 2, 2, 0, 1, 0, 0, 1, 2, 2, 2, 2, 1, 0, 0, 0, 2, 2, 1, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 2, 0, 2, 1, 1, 2, 2, 1, 1, 2, 2, 2, 0, 2, 2, 2, 1, 1, 2, 1, 1, 0, 1, 2, 0, 1, 2, 0, 2, 0, 2, 2, 0, 2, 1, 0, 2, 0, 1, 1, 0, 2, 1, 1, 2, 0, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 0, 0, 2, 2, 1, 2, 0, 1, 1, 2, 1, 1, 1, 1, 0, 1, 0, 2, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 1, 2, 2, 1, 1, 2, 1, 0, 1, 2, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 2, 0, 2, 1, 1, 0, 2, 2, 2, 0, 2, 1, 2, 2, 1, 1, 1, 1, 1, 0, 0, 0, 1, 2, 2, 2, 0, 0, 0, 0, 2, 1, 0, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 1, 0, 2, 1, 1, 1, 0, 2, 1, 2, 0, 1, 1, 1, 0, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 2, 1, 2, 2, 0, 1, 2, 0, 2, 1, 1, 0, 1, 1, 1, 2, 0, 0, 1, 2, 1, 2, 0, 1, 2, 1, 2, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 1, 1, 0, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 2, 2, 0, 2, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 1, 1, 2, 0, 0, 2, 1, 2, 2, 2, 2, 2, 1, 0, 2, 2, 0, 1, 1, 0, 2, 2, 1, 1, 0, 2, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 0, 0, 2, 2, 0, 1, 1, 0, 2, 2, 1, 2, 1, 1, 0, 2, 2, 0, 1, 2, 1, 2, 1, 0, 1, 1, 0, 0, 2, 2, 2, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0, 1, 1, 0, 2, 0, 0, 2, 0, 2, 0, 0, 0, 2, 1, 2, 2, 1, 0, 2, 2, 1, 2, 2, 0, 2, 0, 1, 0, 0, 1, 2, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 2, 2, 2, 1, 0, 2, 2, 2, 0, 2, 1, 2, 1, 2, 0, 1, 1, 0, 0, 0, 2, 2, 1, 1, 1, 2, 0, 2, 1, 1, 1, 0, 0, 1, 2, 2, 0, 0, 0, 1, 1, 2, 1, 2, 0, 2, 1, 2, 2},
		k:    25000000,
	}

	fmt.Println(smallestDistancePair(input.nums, input.k))
}

func smallestDistancePair(nums []int, k int) int {
	diffs := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i != j {
				diff := abs(nums[i] - nums[j])
				diffs = append(diffs, diff)
			}
		}
	}

	sort.Ints(diffs)
	res := diffs[k-1]

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
