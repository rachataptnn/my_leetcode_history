package main

import "fmt"

func main() {
	fmt.Println(scoreOfString("hello"))
}

func scoreOfString(s string) int {
	sum, tmp := 0, 0
	for i := 0; i < len(s)-1; i++ {
		tmp = abs(int(s[i]), int(s[i+1]))
		sum += tmp
	}

	return sum
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
