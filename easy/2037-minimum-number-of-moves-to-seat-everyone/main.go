// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	fmt.Println(minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	move := 0
	for i, _ := range seats {
		move += abs(seats[i] - students[i])
	}

	return move
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
