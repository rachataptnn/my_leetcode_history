package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
	// this is the one of cases	
	// Input: root = [4,2,7,1,3,6,9]
	// Output: [4,7,2,9,6,3,1]
	
	input := []int{1,2}
	tree := arrayToTree(input)
	inv := invertTree(tree)
	fmt.Printf("\n%+v\n", inv)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
	
	finalRevArray := getReverseArrayOfEachDegree(root, root.Val)
	reversedTree := arrayToTree(finalRevArray)

	return reversedTree
}

func getReverseArrayOfEachDegree(root *TreeNode, firstValue int) []int {
	q := Queue{}
	q.Enqueue(root)

	var finalRevArray []int
	finalRevArray = append(finalRevArray, firstValue)

	for !q.IsEmpty() {
		levelSize := q.Size()   // Get the number of nodes in the current level (levelSize).
		temp := make([]*TreeNode, levelSize)
	
		// Traverse levelSize number of nodes and store them in a temporary list or array.
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue()
			temp[i] = node
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		var eachDegreeVal []int
		for _, v := range q.items {
			eachDegreeVal = append(eachDegreeVal, v.Val)
		}

		revArr := reverseArray(eachDegreeVal)
		finalRevArray = append(finalRevArray, revArr...)
	}

	return finalRevArray
}

func reverseArray(input []int) []int {
	reversedArr := []int{}
	for i:=len(input); i>0; i-- {
		reversedArr = append(reversedArr, input[i-1])
	}

	return reversedArr
}


// helper function for convert data structure (array to tree)
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

        if i < len(arr) {
            if arr[i] != 0 {
                node.Left = &TreeNode{Val: arr[i]} // after exec this line, root is updated
                queue = append(queue, node.Left)
            }
            i++
        }

        if i < len(arr) {
            if arr[i] != 0 {
                node.Right = &TreeNode{Val: arr[i]}
                queue = append(queue, node.Right)
            }
            i++
        }
    }

    return root
}

// things for queue data structure
type Queue struct {
    items []*TreeNode
    size  int
}

// Enqueue adds a TreeNode to the end of the queue
func (q *Queue) Enqueue(node *TreeNode) {
    q.items = append(q.items, node)
    q.size++
}

// Dequeue removes and returns the TreeNode from the front of the queue
func (q *Queue) Dequeue() *TreeNode {
    if q.size == 0 {
        return nil
    }
    node := q.items[0]
    q.items = q.items[1:]
    q.size--
    return node
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
    return q.size
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
    return q.size == 0
}