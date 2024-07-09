// https://leetcode.com/problems/average-waiting-time/?envType=daily-question&envId=2024-07-09

package main

import (
	"fmt"
)

func main() {
	// input := [][]int{{1, 2}, {2, 5}, {4, 3}}
	input := [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}

	res := averageWaitingTime(input)
	fmt.Println(res)
}

func averageWaitingTime(customers [][]int) float64 {
	attempt := customers[0][0]
	duration := customers[0][1]
	start := attempt + duration
	sumDuration := duration

	for i := 1; i < len(customers); i++ {
		duration := customers[i][1]
		attemptTime := customers[i][0]

		waitingTime := 0
		if start < attemptTime {
			start = attemptTime + duration
			waitingTime = duration
		} else {
			start += duration
			waitingTime = start - attemptTime
		}

		sumDuration += waitingTime
	}
	return float64(sumDuration) / float64(len(customers))
}
