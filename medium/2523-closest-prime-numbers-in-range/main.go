package main

import "fmt"

func main() {
	// fmt.Println(closestPrimes(10, 19))

	// pass 38/66
	fmt.Println(closestPrimes(19, 31))
}

func closestPrimes(left int, right int) []int {
	if left == right {
		return []int{-1, -1}
	}

	min := 99999999
	firstIndex := -1
	secondIndex := -1

	firstRes := -1
	secondRes := -1

	for i := left; i < right; i++ {
		if isPrime(i) && firstIndex == -1 && secondIndex == -1 {
			firstIndex = i
		} else if isPrime(i) {
			secondIndex = i

			tmpMin := secondIndex - firstIndex
			if tmpMin < min {
				min = tmpMin
				firstRes, secondRes = firstIndex, secondIndex
			}

			firstIndex = i
			secondIndex = -1
		}
	}

	return []int{firstRes, secondRes}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
