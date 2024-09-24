// https://leetcode.com/problems/gas-station/?envType=problem-list-v2&envId=greedy&difficulty=MEDIUM

package main

import "fmt"

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	// 21/39
	// gas := []int{5, 1, 2, 3, 4}
	// cost := []int{4, 4, 1, 5, 1}

	// 25/39
	// gas := []int{4, 5, 2, 6, 5, 3}
	// cost := []int{3, 2, 7, 3, 2, 9}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	gasLeftInTank := 0
	start := 0 // at first, lets assume that the first station(index 0) is the answer

	for i := 0; i < len(gas); i++ {
		gasLeftInTank += gas[i] - cost[i]
		if gasLeftInTank < 0 { // and then just iterate till we found that start with index 0 is not possible because the tank is no gas left(gas<0)
			start = i + 1
			gasLeftInTank = 0 // these 2 lines is just reset, lets assume that the next station is the answer
		}

		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}
	// if we comment this condition, the test case below would got answer = 2
	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}
	// which is wrong, it is not possible to run because
	// sumGas = 9
	// sumCost = 10
	// the gas is not enough

	return start
}
