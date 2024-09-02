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

func leafNodes(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		fmt.Print(root.Val, " ")
	}
	if root.Left != nil {
		leafNodes(root.Left)
	}
	if root.Right != nil {
		leafNodes(root.Right)
	}
}

func main() {

	tree := &TreeNode{Val: 50}
	tree.insert(30)
	tree.insert(20)
	tree.insert(55)
	tree.insert(66)
	tree.insert(99)
	tree.insert(8)
	tree.insert(74)
	tree.insert(24)
	tree.insert(87)
	tree.insert(6)
	tree.insert(52)
	tree.insert(60)

	leafNodes(tree)
}
