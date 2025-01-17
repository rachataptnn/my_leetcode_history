package main

import "fmt"

func main() {
	in := []int{1, 1, 0}
	fmt.Println(doesValidArrayExist(in))
}

func doesValidArrayExist(derived []int) bool {
	n := len(derived)

	// Check with original array starting with 0
	original0 := []int{0}
	for i := 0; i < n; i++ {
		xor := derived[i] ^ original0[i]
		original0 = append(original0, xor)
	}
	lastIndex := len(original0) - 1
	checkForZero := original0[0] == original0[lastIndex]

	// Check with original array starting with 1
	original1 := []int{1}
	for i := 0; i < len(derived); i++ {
		xor := derived[i] ^ original1[i]
		original1 = append(original1, xor)
	}
	lastIndex = len(original1) - 1
	checkForOne := original1[0] == original1[lastIndex]

	return checkForZero || checkForOne
}
