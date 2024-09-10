// https://leetcode.com/problems/find-bottom-left-tree-value/

package main

func main() {
	root := &TreeNode{
		Val: 2,
	}

	root.Left = &TreeNode{
		Val: 2,
	}

	root.Right = &TreeNode{
		Val: 3,
	}

	findBottomLeftValue(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	s := states{}

	s.preOrderTraversal(root, 0)

	return 0
}

type states struct {
	vwl   []valWithLevel
	level int
}

type valWithLevel struct {
	val   int
	level int
}

func (s *states) preOrderTraversal(root *TreeNode, prevNodeVal int) {
	if root == nil {
		if prevNodeVal > 0 {
			s.vwl = append(s.vwl, valWithLevel{
				val:   prevNodeVal,
				level: s.level,
			})
		}

		// s.level -1

		return
	}

	s.level += 1

	s.preOrderTraversal(root.Left, root.Val) // Traverse left subtree
	s.preOrderTraversal(root.Right, 0)       // Traverse right subtree
}
