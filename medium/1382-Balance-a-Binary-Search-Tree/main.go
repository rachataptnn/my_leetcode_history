package main

import (
	"fmt"
)

// TreeNode represents a node in the binary search tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var values []int
	inorder(root, &values)
	return buildBalancedBST(values, 0, len(values)-1)
}

func inorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, values)
	*values = append(*values, node.Val)
	inorder(node.Right, values)
}

func buildBalancedBST(values []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &TreeNode{Val: values[mid]}

	node.Left = buildBalancedBST(values, left, mid-1)
	node.Right = buildBalancedBST(values, mid+1, right)

	return node
}

// Helper function to print the tree in-order
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Print(node.Val, " ")
	printInorder(node.Right)
}

// Helper function to create an unbalanced BST for testing
func createUnbalancedBST() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 4}
	return root
}

func main() {
	// Create an unbalanced BST
	unbalancedRoot := createUnbalancedBST()

	fmt.Println("Unbalanced BST (in-order):")
	printInorder(unbalancedRoot)
	fmt.Println()

	// Balance the BST
	balancedRoot := balanceBST(unbalancedRoot)

	fmt.Println("Balanced BST (in-order):")
	printInorder(balancedRoot)
	fmt.Println()
}
