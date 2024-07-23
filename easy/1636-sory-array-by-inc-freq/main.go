package main

import (
	"fmt"
	"sort"
)

func main() {
	// base case 1
	// nums := []int{1, 1, 2, 2, 2, 3}

	// nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}

	// // normal case: order incorrect1
	// nums := []int{2, 3, 1, 3, 2}

	// 55/180
	// nums := []int{2, -3, 2, 2}

	// 56/180
	// nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}

	// 61/180
	nums := []int{8, -8, 2, -8, -5, -3}

	fmt.Println(frequencySort(nums))
}

func frequencySort(nums []int) []int {
	matrix := prepareFreqMatrix(nums)

	res := []int{}
	for _, v := range matrix {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		res = append(res, v...)
	}

	return res
}

func prepareFreqMatrix(nums []int) [][]int {
	// used for count the frequency
	// key    val
	// 1      10    (found number 1, 10 times)
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	// revert key and val
	// from [1]10, [12]10, [3]10, [8]5
	// to   [10][1, 12, 3], [5][8]
	mapOfArr := make(map[int][]int)
	for k, freq := range freqMap {
		for i := 0; i < freq; i++ {
			mapOfArr[freq] = append(mapOfArr[freq], k)

		}
	}

	// sorting the lowest freq come first!
	// and i call it matrix not 2d array bcuz is lil bit too much cute <3
	keys := make([]int, 0, len(mapOfArr))
	for k := range mapOfArr {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	freqMatrix := make([][]int, len(keys))
	for i, key := range keys {
		freqMatrix[i] = mapOfArr[key]
	}
	return freqMatrix
}
