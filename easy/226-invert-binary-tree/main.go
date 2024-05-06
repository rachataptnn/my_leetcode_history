package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func main() {
    // Input: root = [1,2]
    input := []int{1, 2}
    root := arrayToTree(input)
    invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    invertedRoot := &TreeNode{Val: root.Val}
    invertedRoot.Left = invertTree(root.Right)
    invertedRoot.Right = invertTree(root.Left)
    return invertedRoot
}

func arrayToTree(arr []int) *TreeNode {
    if len(arr) == 0 {
        return nil
    }
    root := &TreeNode{Val: arr[0]}
    queue := []*TreeNode{root}
    i := 1
    for i < len(arr) {
        node := queue[0]
        queue = queue[1:]
        if i < len(arr) && arr[i] != 0 {
            node.Left = &TreeNode{Val: arr[i]}
            queue = append(queue, node.Left)
        } else {
            node.Left = nil
        }
        i++
        if i < len(arr) && arr[i] != 0 {
            node.Right = &TreeNode{Val: arr[i]}
            queue = append(queue, node.Right)
        } else {
            node.Right = nil
        }
        i++
    }
    return root
}