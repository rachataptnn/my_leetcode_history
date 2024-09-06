// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/?envType=daily-question&envId=2024-09-06

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	filteredHead := modifiedList(nums, head)
	printLinkedList(filteredHead)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	skipsMap := make(map[int]bool)
	for _, v := range nums {
		skipsMap[v] = true
	}

	dummy := &ListNode{}
	current := dummy

	for head != nil {
		if !skipsMap[head.Val] {
			current.Next = &ListNode{Val: head.Val}
			current = current.Next
		}
		head = head.Next
	}

	return dummy.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
