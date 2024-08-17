// https://leetcode.com/problems/greatest-common-divisor-of-strings/

package main

import "fmt"

func main() {
	// str1 := "ABCABC"
	// str2 := "ABC"

	// str1 := "ABABAB"
	// str2 := "ABAB"

	// str1 := "LEET"
	// str2 := "CODE"

	// 78/124
	str1 := "ABABABAB"
	str2 := "ABAB"

	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	stringsUnDivisible := str1+str2 != str2+str1
	if stringsUnDivisible {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}

// O(log n)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// huge thanks for Claude <3

// these explanation is DOPE!
// Let's see an example with gcd(48, 18):
// First iteration:  a = 18, b = 48%18 = 12
// Second iteration: a = 12, b = 18%12 = 6
// Third iteration:  a = 6,  b = 12%6  = 0
// Loop ends, return 6

// O(n)
func naiveGCD(a, b int) int {
	smaller := min(a, b)
	for i := smaller; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}
