package main

import (
	"fmt"
	"strconv"
)

func main() {
	// num := "123"
	num := "112"

	fmt.Println(countBalancedPermutations(num))
}

func countBalancedPermutations(num string) int {
	in := []rune(num)
	perms := permute(in, 0, len(in)-1)
	cnt := 0

	for _, v := range perms {
		if isBalanced(v) {
			cnt++
		}
	}

	return cnt
}

func permute(in []rune, l, r int) []string {
	var out []string
	if l == r {
		out = append(out, string(in))
	} else {
		used := make(map[rune]bool)

		for i := l; i <= r; i++ {
			if used[in[i]] {
				continue
			}
			used[in[i]] = true

			in[l], in[i] = in[i], in[l]
			out = append(out, permute(in, l+1, r)...)
			in[l], in[i] = in[i], in[l] // backtrack
		}
	}
	return out
}

func isBalanced(s string) bool {
	sumOdd := 0
	sumEven := 0

	for i, v := range s {
		num, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			sumEven += num
		} else {
			sumOdd += num
		}
	}

	return sumEven == sumOdd
}
