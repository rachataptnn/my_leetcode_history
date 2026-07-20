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
	noMoneyChild := 0
	for v := range children {
		if money >= noMoneyChild && money-8 >= 0 {
			money -= 8
			childGot8++
		}

		noMoneyChild = children - childGot8
		if money < noMoneyChild {
			childGot8--
			noMoneyChild++
			money += 8
		}

		if noMoneyChild*4 == money {
			childGot8--
			break
		}

		fmt.Println(v)
	}

	return childGot8
}
