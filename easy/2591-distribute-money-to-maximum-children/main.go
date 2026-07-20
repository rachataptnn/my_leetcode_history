package main

import "fmt"

func main() {
	money := 20
	children := 3
	result := distMoney(money, children)

	fmt.Println(result)
}

func distMoney(money int, children int) int {
	if money < 8 {
		return -1
	}

	childGot8 := 0
	for v := range children {
		childNo := v + 1
		if money < 8 {
			break
		}
		money -= childNo * 8
		childGot8++
	}
	if money == 0 {
		return childGot8
	}

	noMoneyChild := children - childGot8
	if money >= noMoneyChild {
		return childGot8
	} else {
		return childGot8 - 1
	}

	return 0
}
