// https://leetcode.com/problems/merge-nodes-in-between-zeros/?envType=daily-question&envId=2024-07-04

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	inputList := createList()
	mergeNodes(inputList)
}

func mergeNodes(head *ListNode) *ListNode {
	resultList := mergeEachGap(head, &ListNode{}, 0)
	return resultList
}

func mergeEachGap(inputNodes, outputNodes *ListNode, sum int) *ListNode {
	isListEnd := inputNodes.Next == nil
	if isListEnd {
		outputNodes.Val = sum
		return outputNodes
	}

	switch inputNodes.Val {
	case 0:
		if sum != 0 {
			outputNodes.Val = sum

			sum = 0
			isListEnd := inputNodes.Next == nil
			if !isListEnd {
				outputNodes.Next = &ListNode{}
				outputNodes = outputNodes.Next
			}
		}
		mergeEachGap(inputNodes.Next, outputNodes, sum)
	default:
		sum += inputNodes.Val
		mergeEachGap(inputNodes.Next, outputNodes, sum)
	}

	return outputNodes
}

// Input: l1 = [0,3,1,0,4,5,2,0]
// forgive me for using this way init input :P
func createList() *ListNode {
	list := &ListNode{Val: 0}
	list.Next = &ListNode{Val: 3}
	list.Next.Next = &ListNode{Val: 1}
	list.Next.Next.Next = &ListNode{Val: 0}
	list.Next.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	list.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	list.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}

	return list
}

// somenote: At first i think its toooooooo easy
// yeah im done its easy as i think
// but for sure... this problem provide me better understanding liked list <3
