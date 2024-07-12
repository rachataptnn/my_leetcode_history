// https://leetcode.com/problems/merge-k-sorted-lists/description/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type minWithIndex struct {
	index int
	min   int
}

func main() {
	arrs := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}

	lists := prepareListArr(arrs)
	resultList := mergeKLists(lists)
	printLinkedList(resultList)
}

func mergeKLists(lists []*ListNode) *ListNode {
	resultList := &ListNode{}
	findMinNumInLists(lists, resultList)
	return resultList
}

func findMinNumInLists(lists []*ListNode, resultList *ListNode) *ListNode {
	minCh := make(chan minWithIndex, len(lists))
	defer close(minCh)

	for index, list := range lists {
		min := list.Val
		go findMinNumInList(minCh, list, index, min)
	}

	combinedMin := []minWithIndex{}
	for i := 0; i < len(lists); i++ {
		combinedMin = append(combinedMin, <-minCh)
	}

	theMin := sortAndFindMin(combinedMin)
	resultList.Val = theMin.min
	tempList := &ListNode{}
	resultList.Next = tempList

	if lists[theMin.index].Next != nil {
		lists[theMin.index] = lists[theMin.index].Next
	} else {
		lists = removeList(lists, theMin.index)
	}

	if len(lists) != 0 {
		return findMinNumInLists(lists, resultList.Next)
	} else {
		return resultList
	}
}

func findMinNumInList(minCh chan minWithIndex, list *ListNode, index, min int) {
	if list.Next == nil {
		minCh <- minWithIndex{
			index: index,
			min:   min,
		}
		return
	}

	if list.Val < min {
		min = list.Val
	}
	findMinNumInList(minCh, list.Next, index, min)
}

func sortAndFindMin(arr []minWithIndex) minWithIndex {
	if len(arr) == 0 {
		return minWithIndex{}
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].min > key.min {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr[0]
}

func removeList(lists []*ListNode, index int) []*ListNode {
	if index < 0 || index >= len(lists) {
		return lists // Index out of range, return the original slice
	}
	return append(lists[:index], lists[index+1:]...)
}

// Helper Funcs
func prepareListArr(arrs [][]int) (listArr []*ListNode) {
	for _, arr := range arrs {
		list := arrayToLinkedList(arr)
		listArr = append(listArr, list)
	}

	return listArr
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	return &ListNode{
		Val:  arr[0],
		Next: arrayToLinkedList(arr[1:]),
	}
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
