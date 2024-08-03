package main

import "fmt"

func main() {
	// Construct the binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	distance := 3

	fmt.Printf("\n\nAns%v\n", countPairs(root, distance))
}

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	res := 0
	leaves(root, &res, distance)
	return res
}

func leaves(node *TreeNode, res *int, d int) []int {
	if node.Left == nil && node.Right == nil {
		return []int{1}
	}

	left, right := []int{}, []int{}
	if node.Left != nil {
		left = leaves(node.Left, res, d)
	}
	if node.Right != nil {
		right = leaves(node.Right, res, d)
	}

	for _, a := range left {
		for _, b := range right {
			if a+b <= d {
				*res++
			}
		}
	}

	nodes := append(left, right...)
	for i := range nodes {
		nodes[i]++
	}
	return nodes
}
