// https://leetcode.com/problems/balance-a-binary-search-tree/description/?envType=daily-question&envId=2024-06-26

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	balanceBST(&tn)
}

func balanceBST(root *TreeNode) *TreeNode {
	return nil
}
