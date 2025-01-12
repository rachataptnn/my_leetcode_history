package main

import (
	"fmt"
)

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for l, index := range indices {
		result[index] = s[l]
	}
	return string(result)
}
