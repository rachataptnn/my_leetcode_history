// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/?envType=daily-question&envId=2024-09-06

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	// head = [1,2,3,4,5]

	head := &ListNode{}
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Output: [4,5]
	fmt.Println(modifiedList(nums, head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	filtered := &ListNode{}

	diveIntoLL(head, filtered, nums)

	return nil
}

func diveIntoLL(head, filtered *ListNode, num []int) {
	if head == nil {
		return
	}

	needToAdd := true

	for _, v := range num {
		if head.Val == v {
			needToAdd = false
		}
	}

	if needToAdd {
		filtered.Val = head.Val
		if head.Next != nil {
			filtered.Next = &ListNode{}
			diveIntoLL(head.Next, filtered.Next, num)
		} else {
			return
		}
	}

	diveIntoLL(head.Next, filtered, num)
}
