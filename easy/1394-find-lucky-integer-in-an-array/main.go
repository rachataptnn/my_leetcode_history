package main

import "fmt"

func main() {
	arr := []int{2, 2, 3, 4}
	// arr := []int{1, 2, 2, 3, 3, 3}
	res := findLucky(arr)
	fmt.Println(res)
}

func findLucky(arr []int) int {
	largest := -1
	freqM := make(map[int]int)
	for _, v := range arr {
		freqM[v]++
	}

	for key, val := range freqM {
		if key == val && key > largest {
			largest = key
		}
	}

	return largest
}
