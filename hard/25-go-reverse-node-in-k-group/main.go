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

// start

func reverseKGroup(head *ListNode, k int) *ListNode {
	s := states{
		head: head,
		k:    k,
	}

	return s.reverseProcess()
}

type states struct {
	head *ListNode
	k    int

	tempArr []int
	res     *ListNode
	resTail **ListNode

	stepAmt int
}

func (s *states) reverseProcess() *ListNode {
	if s.head == nil {
		s.appendToResultList()
		return s.res
	}

	s.tempArr = append(s.tempArr, s.head.Val)
	if len(s.tempArr) == s.k {
		reverseSlice(s.tempArr)

		if s.res == nil {
			s.initializeResultList()
		} else {
			s.appendToResultList()
		}
		s.tempArr = []int{}
		s.stepAmt++
	}
	s.head = s.head.Next
	return s.reverseProcess()
}

func (s *states) initializeResultList() {
	s.res = &ListNode{Val: s.tempArr[0]}
	current := s.res
	for _, v := range s.tempArr[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	s.resTail = &current
}

func (s *states) appendToResultList() {
	for _, v := range s.tempArr {
		(*s.resTail).Next = &ListNode{Val: v}
		*s.resTail = (*s.resTail).Next
	}
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
