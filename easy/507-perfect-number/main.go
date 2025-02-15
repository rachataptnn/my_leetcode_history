package main

import "fmt"

func main() {
	// pass 97/98
	num := 2732
	fmt.Println(checkPerfectNumber(num))
}

// func checkPerfectNumber(num int) bool {
// 	if num%2 != 0 {
// 		return false
// 	}

// 	sum := 0
// 	for i := 1; i < num; i++ {
// 		if num%i == 0 {
// 			sum += i
// 		}
// 	}

// 	return sum == num
// }

func checkPerfectNumber(num int) bool {
	if num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336 {
		return true
	}
	return false
}

// i love this solution
// https://leetcode.com/problems/perfect-number/solutions/4235608/there-are-only-five-perfect-numbers-below-10-8/
