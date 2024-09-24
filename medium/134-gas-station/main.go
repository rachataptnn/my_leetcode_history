// https://leetcode.com/problems/gas-station/?envType=problem-list-v2&envId=greedy&difficulty=MEDIUM

package main

import "fmt"

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	s := states{
		useGas:  cost,
		fillGas: gas,
	}

	for currentStation := 0; currentStation < len(gas); currentStation++ {
		s.gasLeft = s.fillGas[currentStation]
		if s.isCircuitPossible(currentStation) {
			return currentStation
		}
	}

	return -1
}

type states struct {
	gasLeft int
	useGas  []int
	fillGas []int
}

func (s *states) isCircuitPossible(currentStation int) bool {
	count := 0
	for count < len(s.useGas) {
		canGoNext := s.gasLeft-s.useGas[currentStation] > 0
		if canGoNext {
			nextStation := currentStation + 1
			if nextStation >= len(s.useGas) {
				nextStation = 0
			}

			s.gasLeft = s.gasLeft - s.useGas[currentStation] + s.fillGas[nextStation]
			count++
		} else {
			return false
		}
	}

	return true
}
