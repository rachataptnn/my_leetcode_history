// https://leetcode.com/problems/n-ary-tree-postorder-traversal/?envType=daily-question&envId=2024-08-26

package main

import "fmt"

func main() {
	root := Node{}

	res := postorder(&root) // Post-order traversal function from earlier
	fmt.Println("Post-order Traversal:", res)
}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	postOrderTraversalNary(root, &res)

	return res
}

func postOrderTraversalNary(root *Node, result *[]int) {
	if root == nil {
		return
	}

	// Visit all children
	for _, child := range root.Children {
		postOrderTraversalNary(child, result) // Recursive call for each child
	}

	*result = append(*result, root.Val)
}
