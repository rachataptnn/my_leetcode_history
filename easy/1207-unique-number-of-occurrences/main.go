package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

func uniqueOccurrences(arr []int) bool {
	freqM := make(map[int]int)
	for _, v := range arr {
		freqM[v]++
	}

	freqM2 := make(map[int]int)
	for _, val := range freqM {
		freqM2[val]++
		if freqM2[val] > 1 {
			return false
		}
	}

	return true
}
