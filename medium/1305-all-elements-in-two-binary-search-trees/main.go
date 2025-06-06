package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// root1 := []int{2, 1, 4}
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}

	// root2 := []int{1, 0, 3}
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 0}
	root2.Right = &TreeNode{Val: 3}

	getAllElements(root1, root2)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var result1 []int
	var result2 []int
	inorderTraversal(root1, &result1)
	inorderTraversal(root2, &result2)

	res := append(result1, result2...)
	sort.Ints(res)

	return res
}

func inorderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, result)  // visit left subtree
	*result = append(*result, root.Val)  // visit root
	inorderTraversal(root.Right, result) // visit right subtree
}
