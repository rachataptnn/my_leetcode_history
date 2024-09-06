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
	dummy := &ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {
		if contains(nums, current.Val) {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return dummy.Next
}

func contains(nums []int, val int) bool {
	for _, v := range nums {
		if v == val {
			return true
		}
	}
	return false
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
