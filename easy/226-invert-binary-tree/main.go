package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
	// this is the one of cases	
	// Input: root = [4,2,7,1,3,6,9]
	// Output: [4,7,2,9,6,3,1]
	
	input := []int{4,2,7,1,3,6,9}
	tree := convertArrayToTree(input)
	invertTree(tree)
}

func invertTree(root *TreeNode) *TreeNode {
	return nil
}

// helper function
func convertArrayToTree(input []int) *TreeNode {
    if len(input) == 0 {
        return nil
    }

	// the first degree always have 1 node
	root := &TreeNode{Val: input[0]}

	// second degree must be the index 1,2
	root.Left = &TreeNode{Val: input[1]}        // is this i*2+1    lets say i=0 at second degree
    root.Right = &TreeNode{Val: input[2]}       // i*2 + 2

	// third degree must be the index 3,4
	root.Left.Left = &TreeNode{Val: input[3]}   // i=1    i*2+1 = 3
    root.Left.Right = &TreeNode{Val: input[4]}  //        i*2+2 = 4
	// and 5,6
    root.Right.Left = &TreeNode{Val: input[5]}  // i=2    i*2+1 = 5
    root.Right.Right = &TreeNode{Val: input[6]} //		  i*2+2 = 6   
	// ...and... at the 3rd degree, node amount is 2x from prev degree
    
	// so... i think i is the outer index 
	// should i declare variable degreeSize for each degree?

	return root
}
