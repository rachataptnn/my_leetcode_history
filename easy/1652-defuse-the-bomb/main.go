// https://leetcode.com/problems/defuse-the-bomb/

package main

import "fmt"

func main() {
	// code := []int{5, 7, 1, 4}
	// k := 3

	code := []int{2, 4, 9, 3}
	k := -2
	fmt.Println(decrypt(code, k))
}

func decrypt(code []int, k int) []int {
	doublingCode := append(code, code...)
	if k > 0 {
		res := kmorethanzero(code, doublingCode, k)
		return res
	} else if k < 0 {
		res := klessthanzero(code, doublingCode, k)
		return res
	} else {
		res := make([]int, len(code))
		return res
	}
}

func kmorethanzero(code, doublingCode []int, k int) []int {
	res := []int{}
	cnt := 0
	sum := 0
	for i := 0; i < len(code); i++ {
		for j := i; j < len(doublingCode); j++ {
			sum += doublingCode[j+1]
			cnt++
			if cnt == k {
				res = append(res, sum)
				sum = 0
				cnt = 0
				break
			}
		}
	}

	return res
}

func klessthanzero(code, doublingCode []int, k int) []int {
	k = k * -1

	res := []int{}
	cnt := 0
	sum := 0

	for i := len(code); i < len(doublingCode); i++ {
		for j := i; j >= 0; j-- {
			sum += doublingCode[j-1]
			cnt++
			if cnt == k {
				res = append(res, sum)
				sum = 0
				cnt = 0
				break
			}
		}
	}

	return res
}
