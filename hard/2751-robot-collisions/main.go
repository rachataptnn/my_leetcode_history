// https://leetcode.com/problems/robot-collisions/description/?envType=daily-question&envId=2024-07-13

package main

import (
	"fmt"
	"sort"
)

type input struct {
	positions  []int
	healths    []int
	directions string
}

func main() {
	// input := input{
	// 	positions:  []int{5, 4, 3, 2, 1},
	// 	healths:    []int{2, 17, 9, 15, 10},
	// 	directions: "RRRRR",
	// }
	// input := input{
	// 	positions:  []int{3, 5, 2, 6},
	// 	healths:    []int{10, 10, 15, 12},
	// 	directions: "RLRL",
	// }
	// input := input{
	// 	positions:  []int{1, 2, 5, 6},
	// 	healths:    []int{10, 10, 11, 11},
	// 	directions: "RLRL",
	// }

	// case 34/2433 output limit exceed
	// input := input{
	// 	positions:  []int{3, 47},
	// 	healths:    []int{46, 26},
	// 	directions: "LR",
	// }

	// case 34/2433 output limit exceed
	// input := input{
	// 	positions:  []int{2, 19, 46},
	// 	healths:    []int{42, 45, 2},
	// 	directions: "LRL",
	// }

	// case 141: เราไป swap original seq มาตลอดเลย5555
	input := input{
		positions:  []int{5, 46, 12},
		healths:    []int{3, 27, 43},
		directions: "RLL",
	}

	res := survivedRobotsHealths(input.positions, input.healths, input.directions)
	fmt.Println(res)
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	robots := make(map[int][]int)
	for i, position := range positions {
		direction := 1
		if directions[i] == 'L' {
			direction = -1
		}
		robots[position] = []int{healths[i], direction, i}
	}

	sortedPositions := append([]int{}, positions...)
	sort.Ints(sortedPositions)
	stack := [][]int{}
	for _, position := range sortedPositions {
		stack = append(stack, robots[position])
		i := len(stack) - 1
		for i-1 >= 0 && stack[i][1] == -1 && stack[i-1][1] == 1 {
			right := stack[i-1]
			left := stack[i]
			survive := []int{}
			stack = stack[:i-1]
			if left[0] > right[0] {
				survive = left
			} else if left[0] < right[0] {
				survive = right
			} else {
				i -= 2
				continue
			}
			survive[0]--
			stack = append(stack, survive)
			i--
		}
	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i][2] < stack[j][2]
	})

	result := []int{}
	for _, robot := range stack {
		result = append(result, robot[0])
	}
	return result
}
