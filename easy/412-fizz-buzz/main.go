// https://leetcode.com/problems/fizz-buzz/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			iStr := strconv.Itoa(i)
			res = append(res, iStr)
		}
	}

	return res
}
