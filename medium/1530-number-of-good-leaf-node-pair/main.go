package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Construct the binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// Find the distance between nodes 4 and 5
	p := root.Left.Left  // Node with value 4
	q := root.Left.Right // Node with value 5

	distance := findDistance(root, p, q)
	fmt.Println("Distance between nodes 4 and 5 is:", distance)
}

func countPairs(root *TreeNode, distance int) int {
	return 0
}

func findDistance(root, p, q *TreeNode) int {
	lca := findLCA(root, p, q)
	return distanceFromLCA(lca, p, 0) + distanceFromLCA(lca, q, 0)
}

func findLCA(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := findLCA(root.Left, p, q)
	right := findLCA(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func distanceFromLCA(lca, node *TreeNode, dist int) int {
	if lca == nil {
		return -1
	}
	if lca == node {
		return dist
	}

	leftDist := distanceFromLCA(lca.Left, node, dist+1)
	if leftDist != -1 {
		return leftDist
	}

	return distanceFromLCA(lca.Right, node, dist+1)
}
