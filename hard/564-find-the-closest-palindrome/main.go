// https://leetcode.com/problems/find-the-closest-palindrome/description/?envType=daily-question&envId=2024-08-24

// fantastic solution: https://leetcode.com/problems/find-the-closest-palindrome/solutions/5675172/o-1-beats-100-0-ms-c-java-python-go-rust-javascript/?envType=daily-question&envId=2024-08-24

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := "835868090839964076"

	// n := "807045053224792883"
	//    807045052250540708 expected
	//    807045053350540708 result
	//    0123456789
	fmt.Println(nearestPalindromic(n))
}

func nearestPalindromic(numberStr string) string {
	number, _ := strconv.ParseInt(numberStr, 10, 64)
	if number <= 10 {
		return strconv.FormatInt(number-1, 10)
	}
	if number == 11 {
		return "9"
	}

	length := len(numberStr)
	leftHalf, _ := strconv.ParseInt(numberStr[:(length+1)/2], 10, 64)

	palindromeCandidates := make([]int64, 5)
	palindromeCandidates[0] = generatePalindromeFromLeft(leftHalf-1, length%2 == 0)
	palindromeCandidates[1] = generatePalindromeFromLeft(leftHalf, length%2 == 0)
	palindromeCandidates[2] = generatePalindromeFromLeft(leftHalf+1, length%2 == 0)
	palindromeCandidates[3] = int64(math.Pow10(length-1)) - 1
	palindromeCandidates[4] = int64(math.Pow10(length)) + 1

	var nearestPalindrome int64
	minDifference := int64(math.MaxInt64)

	for _, candidate := range palindromeCandidates {
		if candidate == number {
			continue
		}
		difference := abs(candidate - number)
		if difference < minDifference || (difference == minDifference && candidate < nearestPalindrome) {
			minDifference = difference
			nearestPalindrome = candidate
		}
	}

	return strconv.FormatInt(nearestPalindrome, 10)
}

func generatePalindromeFromLeft(leftHalf int64, isEvenLength bool) int64 {
	palindrome := leftHalf
	if !isEvenLength {
		leftHalf /= 10
		fmt.Println("")
	}
	for leftHalf > 0 { // i need to learn this part more, maybe in the math section of the solution above
		palindrome = palindrome*10 + leftHalf%10
		leftHalf /= 10
		fmt.Println("")
	}
	return palindrome
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
