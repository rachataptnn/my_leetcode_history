// https://leetcode.com/problems/climbing-stairs/?envType=study-plan-v2&envId=dynamic-programming

package main

import "fmt"

func main() {
	n := 4
	fmt.Println(climbStairs(n))
}

var mem map[int]int = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if val, ok := mem[n]; ok {
		return val
	}

	res := climbStairs(n-1) + climbStairs(n-2)
	mem[n] = res
	return res
}
