package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) insert(val int) {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right.insert(val)
		}
	} else if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Left.insert(val)
		}
	}
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}

func sumOfAllNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := sumOfAllNodes(root.Left)
	rightSum := sumOfAllNodes(root.Right)

	return leftSum + rightSum + root.Val
}

func main() {
	tree := &TreeNode{Val: 50}
	tree.insert(30)
	tree.insert(57)
	tree.insert(67)
	tree.insert(32)
	tree.insert(22)
	tree.insert(45)
	tree.insert(78)

	fmt.Println(sumOfAllNodes(tree))
}