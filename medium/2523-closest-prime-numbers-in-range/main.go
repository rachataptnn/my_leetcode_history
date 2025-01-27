package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(closestPrimes(10, 19))

	// pass 38/66
	// fmt.Println(closestPrimes(19, 31))

	// pass 65/66
	// fmt.Println(closestPrimes(850350, 851803))

	// pass 66/66
	fmt.Println(closestPrimes(850350, 851803))
}

func closestPrimes(left int, right int) []int {
	if left == right {
		return []int{-1, -1}
	}

	min := math.MaxInt64
	firstIndex := -1
	secondIndex := -1

	firstRes := -1
	secondRes := -1

	if right < 30000 {
		for i := left; i <= right; i++ {
			if isPrime(i) {
				if firstIndex == -1 && secondIndex == -1 {
					firstIndex = i
				} else {
					secondIndex = i

					tmpMin := secondIndex - firstIndex
					if tmpMin < min {
						min = tmpMin
						firstRes, secondRes = firstIndex, secondIndex
						if tmpMin == 1 {
							return []int{firstRes, secondRes}
						}
					}

					firstIndex = i
					secondIndex = -1
				}
			}
		}

		return []int{firstRes, secondRes}
	} else {
		isPrime := sieveOfEratosthenes(right)

		for i := left; i <= right; i++ {
			if isPrime[i] && firstIndex == -1 {
				firstIndex = i
			} else if isPrime[i] {
				secondIndex = i

				tmpMin := secondIndex - firstIndex
				if tmpMin < min {
					min = tmpMin
					firstRes, secondRes = firstIndex, secondIndex
					if tmpMin == 1 {
						return []int{firstRes, secondRes}
					}
				}

				firstIndex = i
				secondIndex = -1
			}
		}

		if firstRes == -1 || secondRes == -1 {
			return []int{-1, -1}
		}

		return []int{firstRes, secondRes}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func sieveOfEratosthenes(limit int) []bool {
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}
