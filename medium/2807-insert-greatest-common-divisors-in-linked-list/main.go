//

package main

import "fmt"

func main() {
	head := &ListNode{Val: 18}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 3}

	fmt.Println(insertGreatestCommonDivisors(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	newNode := &ListNode{
		Val:  gcd(head.Val, head.Next.Val),
		Next: head.Next,
	}

	head.Next = newNode

	newNode.Next = insertGreatestCommonDivisors(head.Next.Next)

	return head
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
