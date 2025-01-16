package main

import "fmt"

func main() {
	s := "IDID"
	fmt.Println(diStringMatch(s))
}

func diStringMatch(s string) []int {
	low, high := 0, len(s)
	res := []int{}

	for _, v := range s {
		if v == 'I' {
			res = append(res, low)
			low++
		} else {
			res = append(res, high)
			high--
		}
	}

	res = append(res, low)

	return res
}
