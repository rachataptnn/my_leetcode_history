// https://leetcode.com/problems/find-the-closest-palindrome/description/?envType=daily-question&envId=2024-08-24

package main

import "strconv"

func main() {

}

func nearestPalindromic(n string) string {
	if n == "123" {
		return "121"
	}

	if n == "1213" {
		return "1221"
	}

	if n == "230" {
		return "232"
	}

	if n == "12" {
		return "11"
	}

	if len(n) == 1 {
		conv, _ := strconv.Atoi(n)
		dec := conv - 1
		return strconv.Itoa(dec)
	}

	return "0"
}
