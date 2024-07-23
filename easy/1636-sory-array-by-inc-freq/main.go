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
	matrix := makeFreqMap(nums)
	// sortedMap := sortMapByKey(mapOfArr)

	res := []int{}
	for _, v := range matrix {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		// sort.Ints(v)
		res = append(res, v...)
	}

	return res
}

func makeFreqMap(nums []int) [][]int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	mapOfArr := make(map[int][]int)
	for k, freq := range freqMap {
		for i := 0; i < freq; i++ {
			mapOfArr[freq] = append(mapOfArr[freq], k)

		}
	}

	keys := make([]int, 0, len(mapOfArr))
	for k := range mapOfArr {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	bruh := make([][]int, len(keys))

	// Populate the 2D slice
	for i, key := range keys {
		bruh[i] = mapOfArr[key]
	}

	return bruh
}
