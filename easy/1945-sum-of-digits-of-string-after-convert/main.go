// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/?envType=daily-question&envId=2024-09-03

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "iiii"
	k := 1
	fmt.Println(getLucky(s, k))
}

func getLucky(s string, k int) int {
	alphabetMap := map[rune]string{
		'a': "1", 'b': "2", 'c': "3", 'd': "4", 'e': "5", 'f': "6", 'g': "7", 'h': "8",
		'i': "9", 'j': "10", 'k': "11", 'l': "12", 'm': "13", 'n': "14", 'o': "15",
		'p': "16", 'q': "17", 'r': "18", 's': "19", 't': "20", 'u': "21", 'v': "22",
		'w': "23", 'x': "24", 'y': "25", 'z': "26",
	}

	mapped := ""
	for _, c := range s {
		mapped += alphabetMap[c]
	}

	res := sumMapped(mapped, k)

	return res
}

func sumMapped(mapped string, k int) (res int) {
	sum := 0
	for _, c := range mapped {
		cInt, _ := strconv.Atoi(string(c))
		sum += cInt
	}

	k--
	if k == 0 {
		return sum
	}

	sumStr := strconv.Itoa(sum)
	return sumMapped(sumStr, k)
}
