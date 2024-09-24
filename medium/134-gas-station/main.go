// https://leetcode.com/problems/gas-station/?envType=problem-list-v2&envId=greedy&difficulty=MEDIUM

package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}

	// 21/39
	// gas := []int{5, 1, 2, 3, 4}
	// cost := []int{4, 4, 1, 5, 1}

	// 25/39
	// gas := []int{4, 5, 2, 6, 5, 3}
	// cost := []int{3, 2, 7, 3, 2, 9}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	s := states{
		useGas:       cost,
		fillGas:      gas,
		totalStation: len(gas),
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
	gasLeft      int
	useGas       []int
	fillGas      []int
	totalStation int
}

func (s *states) isCircuitPossible(firstStation int) bool {
	checkInCount := 0

	if firstStation == 3 {
		fmt.Println("")
	}

	for checkInCount != s.totalStation {
		station := firstStation + checkInCount
		if station >= s.totalStation {
			station = abs(s.totalStation, station)
		}

		useGas := s.useGas[station]
		s.gasLeft -= useGas
		if s.gasLeft >= 0 {
			nextStation := station + 1
			if nextStation >= s.totalStation {
				nextStation = 0
			}

			recievedGas := s.fillGas[nextStation]
			s.gasLeft += recievedGas
			checkInCount++
		} else {
			return false
		}
	}

	return true
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
