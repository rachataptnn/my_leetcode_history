// https://leetcode.com/problems/reverse-nodes-in-k-group/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2

	head := initLinkedList(arr)
	fmt.Println(reverseKGroup(head, k))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	s := states{
		head:    head,
		tempArr: []int{},
		k:       k,
	}

	return s.recursiveFn()
}

type states struct {
	head    *ListNode
	res     *ListNode
	resTail **ListNode
	k       int
	tempArr []int
}

func (s *states) recursiveFn() *ListNode {
	if s.head == nil {
		return s.res
	}

	s.tempArr = append(s.tempArr, s.head.Val)
	if len(s.tempArr) == s.k {
		reverseSlice(s.tempArr)
		if s.res == nil {
			s.res = &ListNode{Val: s.tempArr[0]}
			current := s.res
			for _, v := range s.tempArr[1:] {
				current.Next = &ListNode{Val: v}
				current = current.Next
			}
			*(s.resTail) = current
		} else {
			for _, v := range s.tempArr {
				(*s.resTail).Next = &ListNode{Val: v}
				*(s.resTail) = (*s.resTail).Next
			}
		}
		s.tempArr = []int{}
	}

	return s.recursiveFn()
}

func reverseSlice(slice []int) {
	i := 0
	j := len(slice) - 1

	// swapping elements til pointers meet in the middle
	for i < j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
}

// helper function

func initLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for _, v := range arr[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return head
}
