// https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/?envType=daily-question&envId=2024-09-02

package main

import "fmt"

func main() {
	// chalk := []int{5, 1, 5}
	// k := 22

	chalk := []int{3, 4, 1, 2}
	k := 25

	fmt.Println(chalkReplacer(chalk, k))
}

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)

	index := 0
	for k >= 0 {
		k -= chalk[index]

		isChalkEnough := k > 0
		if !isChalkEnough {
			return index
		}

		index++
		if index == n {
			index = 0
		}
	}

	return index
}
