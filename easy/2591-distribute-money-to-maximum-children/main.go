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
		if money >= children {
			return 0
		}
		return -1
	}

	if money < children {
		return -1
	}

	if money/8 == children {
		return children
	}

	childGot8 := 0
	for v := range children {
		childNo := v + 1
		if money <= 8 {
			if money == 4 {
				childGot8--
			}
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
	}
	return childGot8 - 1
}
