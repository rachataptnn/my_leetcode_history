package main

import "fmt"

func main() {
	// Example tree:
	//      2
	//     / \
	//    2   3
	root := &TreeNode{
		Val: 2,
	}

	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 4,
	}

	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Left = &TreeNode{
		Val: 5,
	}
	root.Right.Right = &TreeNode{
		Val: 6,
	}

	root.Right.Left.Left = &TreeNode{
		Val: 7,
	}

	result := findBottomLeftValue(root)
	fmt.Println("Bottom-left value:", result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	s := &states{}
	s.preOrderTraversal(root)

	if len(s.vwl) == 0 {
		return root.Val
	}

	maxLevel := -1
	leftMostVal := 0
	for _, vwl := range s.vwl {
		if vwl.level > maxLevel {
			maxLevel = vwl.level
			leftMostVal = vwl.val
		}
	}

	return leftMostVal
}

type states struct {
	vwl   []valWithLevel
	level int
}

type valWithLevel struct {
	val   int
	level int
}

func (s *states) preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	isLeaf := root.Left == nil && root.Right == nil
	if isLeaf {
		s.vwl = append(s.vwl, valWithLevel{
			val:   root.Val,
			level: s.level,
		})
		return
	}

	s.level++
	s.preOrderTraversal(root.Left)
	s.preOrderTraversal(root.Right)
	s.level--
}
