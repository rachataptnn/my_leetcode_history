package main

import "fmt"

func main() {
	s := "aaabaaaa"
	fmt.Println("res: ", makeFancyString(s))
}

func makeFancyString(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	res := make([]rune, n)
	currCharCnt := 1
	for i, v := range s {
		if i == 0 {
			continue
		}
		if s[i] == s[i-1] {
			currCharCnt++
			if currCharCnt < 3 {
				res = append(res, v)
			}
		} else {
			currCharCnt = 1
		}
	}

	return string(res)
}
