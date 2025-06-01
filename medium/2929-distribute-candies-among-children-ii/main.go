package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(5, 2)) // Sample test
}

// Main function to count valid candy distributions
func distributeCandies(n int, limit int) int64 {
	N := int64(n)
	L := int64(limit)

	totalWays := combinationsWith3Vars(N)
	overLimit1 := countOverLimit(1, N, L)
	overLimit2 := countOverLimit(2, N, L)
	overLimit3 := countOverLimit(3, N, L)

	return totalWays - 3*overLimit1 + 3*overLimit2 - overLimit3
}

// Returns C(n + 2, 2): total ways to split n candies into 3 variables
func combinationsWith3Vars(n int64) int64 {
	return (n + 2) * (n + 1) / 2
}

// PIE helper: returns C2(n - k * limit + adjustment)
func countOverLimit(k int, n, limit int64) int64 {
	var x int64
	switch k {
	case 1:
		x = n - limit + 1
	case 2:
		x = n - 2*limit
	case 3:
		x = n - 3*limit - 1
	}
	return C2(x)
}

// Returns C(n, 2) = n * (n - 1) / 2 (if n >= 2)
func C2(x int64) int64 {
	if x >= 2 {
		return x * (x - 1) / 2
	}
	return 0
}
