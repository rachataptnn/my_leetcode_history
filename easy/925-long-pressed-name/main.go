package main

import "fmt"

func main() {
	// 1 ex true
	// name := "alex"
	// typed := "aaleex"

	// 2 ex false
	name := "saeed"
	typed := "ssaaedd"

	// 76/95 ex false
	// name := "rick"
	// typed := "kric"
	fmt.Println(isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	typedStack := Stack{}

	for i := 0; i < len(typed); i++ {
		if i == 0 {
			typedStack.Push(typed[i])
			continue
		}
		if typed[i] != typed[i-1] {
			typedStack.Push(typed[i])
		}
	}

	return true
}

type Stack struct {
	elements []byte
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value byte) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (byte, bool) {
	if len(s.elements) == 0 {
		return 0, false // stack is empty
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// IsEmpty checks if the stack is empty
// func (s *Stack) IsEmpty() bool {
// 	return len(s.elements) == 0
// }

// Size returns the number of elements in the stack
// func (s *Stack) Size() int {
// 	return len(s.elements)
// }

// Peek returns the top element without removing it
// func (s *Stack) Peek() (int, bool) {
// 	if len(s.elements) == 0 {
// 		return 0, false
// 	}
// 	return s.elements[len(s.elements)-1], true
// }
