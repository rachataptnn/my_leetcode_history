package main

import "fmt"

func main() {
	s := "annabelle"
	k := 2

	fmt.Println(canConstruct(s, k))
}

func canConstruct(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}

	letterFreqMap := make(map[string]int)
	for _, letter := range s {
		letterFreqMap[string(letter)]++
	}

	oddFound := 0
	for _, value := range letterFreqMap {
		if value%2 != 0 {
			oddFound++
		}
	}

	if oddFound > k {
		return false
	}

	return true
}

// O(n) but 0ms lol, why: https://chatgpt.com/c/678233d1-d82c-800d-92fd-286e3e55edb9
// func canConstruct(s string, k int) bool {
//     if len(s) < k {
//         return false
//     }
// 	dp := make([]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		dp[s[i]-'a']++
// 	}
// 	countOdd := 0
// 	for _, v := range dp {
// 		countOdd += v % 2
// 	}
// 	if k < countOdd {
// 		return false
// 	}
// 	return true
// }
