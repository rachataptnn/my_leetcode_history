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
	start := customers[0][0] + customers[0][1]
	sumDuration := customers[0][1]

	for i := 1; i < len(customers); i++ {
		waitingTime := 0
		if start < customers[i][0] {
			start = customers[i][0] + customers[i][1]
			waitingTime = customers[i][1]
		} else {
			start += customers[i][1]
			waitingTime = start - customers[i][0]
		}

		sumDuration += waitingTime
	}
	return float64(sumDuration) / float64(len(customers))
}

func someFn(start, attemptTime, waitingTime, sumDuration int) float64 {

	return 0
}
