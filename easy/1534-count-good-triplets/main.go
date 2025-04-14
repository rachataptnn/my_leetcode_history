package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3

	fmt.Println(
		"res: ", countGoodTriplets(arr, a, b, c),
	)

}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[j]-arr[k]) <= b {
						if abs(arr[i]-arr[k]) <= c {
							res++
						}
					}
				}
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
