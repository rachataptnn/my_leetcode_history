// https://leetcode.com/problems/lemonade-change/?envType=daily-question&envId=2024-08-15

package main

import "fmt"

func main() {
	// bills := []int{5, 5, 5, 10, 20}
	// bills := []int{5, 5, 10, 10, 20}
	bills := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	s := wallet{}

	for _, bill := range bills {
		switch bill {
		case 5:
			s.bill5Cnt++
		case 10:
			if !s.canChange10() {
				return false
			}
		case 20:
			if !s.canChange20() {
				return false
			}
		}
	}

	return true
}

type wallet struct {
	bill5Cnt  int
	bill10Cnt int
}

func (s *wallet) canChange10() bool {
	if s.bill5Cnt > 0 {
		s.bill5Cnt--
		s.bill10Cnt++
		return true
	}

	return false
}

func (s *wallet) canChange20() bool {
	if s.bill10Cnt > 0 && s.bill5Cnt > 0 {
		s.bill5Cnt -= 1
		s.bill10Cnt -= 1
		return true
	}

	if s.bill5Cnt > 2 {
		s.bill5Cnt -= 3
		return true
	}

	return false
}
