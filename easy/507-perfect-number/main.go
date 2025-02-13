package main

import "fmt"

func main() {
	// pass 97/98
	num := 2732
	fmt.Println(checkPerfectNumber(num))
}

func checkPerfectNumber(num int) bool {
	if num%2 != 0 {
		return false
	}

	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}
