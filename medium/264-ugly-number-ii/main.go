// https://leetcode.com/problems/ugly-number-ii/?envType=daily-question&envId=2024-08-18

package main

import (
	"fmt"
)

func main() {
	n := 15
	fmt.Println(nthUglyNumber(n))
}

func nthUglyNumber(n int) int {
	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1

	index2, index3, index5 := 0, 0, 0

	for i := 1; i < n; i++ { // ------------- i = 1 | 2 | 3 | 4 | 5  | 6  | 7  | 8  | 9  | 10  | 11  | 12  | 13  | 14  | 15    | ...
		nextUglyBy2 := uglyNumbers[index2] * 2 // 2 | 4 |   | 6 |    | 8  | 10 |    | 12 | 16  |     | 18  | 20  | 24  | i = n |
		nextUglyBy3 := uglyNumbers[index3] * 3 // 3 |   | 6 |	|    | 9  |    | 12 |    | 15  | 18  |     | 24  |     | i = n |
		nextUglyBy5 := uglyNumbers[index5] * 5 // 5 |   |   |   | 10 |    |    |    | 15 |     | 20  |     |     | 25  | i = n |
		// - minOf(nu2, nu3, nu5) = uglyNumbers = 2 , 3 , 4 , 5 , 6  , 8  , 9  , 10 , 12 , 15  , 16  , 18  , 20  , 24  , ...

		nextUgly := min(nextUglyBy2, min(nextUglyBy3, nextUglyBy5))
		uglyNumbers[i] = nextUgly

		// Increment the index for the multiple that matched the smallest
		if nextUgly == nextUglyBy2 {
			index2++
		}
		if nextUgly == nextUglyBy3 {
			index3++
		}
		if nextUgly == nextUglyBy5 {
			index5++
		}
	}

	return uglyNumbers[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
