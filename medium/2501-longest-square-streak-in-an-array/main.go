// https://leetcode.com/problems/longest-square-streak-in-an-array/description/?envType=daily-question&envId=2024-10-28\

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{4, 3, 6, 16, 8, 2}
	fmt.Println(longestSquareStreak(nums))
}

func longestSquareStreak(nums []int) int {
	mp := make(map[int]int)
	sort.Ints(nums)
	res := -1

	for _, num := range nums {
		sqrt := int(math.Sqrt(float64(num)))

		if sqrt*sqrt == num {
			if val, exists := mp[sqrt]; exists {
				mp[num] = val + 1
				if mp[num] > res {
					res = mp[num]
				}
			} else {
				mp[num] = 1
			}
		} else {
			mp[num] = 1
		}
	}

	return res
}
