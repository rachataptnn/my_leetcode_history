// https://leetcode.com/problems/binary-tree-postorder-traversal/?envType=daily-question&envId=2024-08-25

package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []interface{}{1, nil, 2, 3}

	root := createBinaryTree(arr)
	fmt.Println("Preorder Traversal of the created tree:")

	res := []int{}
	postOrderTraversal(root, &res)

	fmt.Println(res)
}

func postOrderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left, result)
	postOrderTraversal(root.Right, result)
	*result = append(*result, root.Val)
}

// Helper functions
func createBinaryTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < len(arr) {
		current := queue[0]
		queue = queue[1:]

		// Create the left child
		if index < len(arr) && arr[index] != nil {
			current.Left = &TreeNode{Val: arr[index].(int)}
			queue = append(queue, current.Left)
		}
		index++

		// Create the right child
		if index < len(arr) && arr[index] != nil {
			current.Right = &TreeNode{Val: arr[index].(int)}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

// Helper function to print the tree (Preorder traversal)
func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}
