// https://leetcode.com/problems/valid-parentheses/description/

package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	type1 := 0
	type2 := 0
	type3 := 0

	isCloseBeforeOpen := false

	stack := Stack{}

	for _, v := range s {
		if v == '(' {
			type1++
			stack.Push('(')
		} else if v == '[' {
			type2++
			stack.Push('[')
		} else if v == '{' {
			type3++
			stack.Push('{')
		} else if v == ')' {
			openBracket := stack.Pop()
			if openBracket != '(' {
				return false
			}

			type1--
			isCloseBeforeOpen = type1 < 0
		} else if v == ']' {
			openBracket := stack.Pop()
			if openBracket != '[' {
				return false
			}

			type2--
			isCloseBeforeOpen = type2 < 0
		} else if v == '}' {
			openBracket := stack.Pop()
			if openBracket != '{' {
				return false
			}

			type3--
			isCloseBeforeOpen = type3 < 0
		}

		if isCloseBeforeOpen {
			return false
		}
	}

	isPaired := type1 == 0 && type2 == 0 && type3 == 0

	return isPaired
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty!")
		return -1
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
