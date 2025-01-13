package main

import "fmt"

func main() {
	ex1 := "abaacbcbb" // 5
	fmt.Println(minimumLength(ex1))
}

func minimumLength(s string) int {
	m := make(map[rune]int)

	cnt := 0
	for _, v := range s {
		m[v]++
	}

	for _, v := range m {
		if v%2 == 0 {
			cnt += 2
		} else {
			cnt += 1
		}
	}

	return cnt
}
