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
	fmt.Print(" ", root.Val)
	inOrder(root.Right)
}

func isBalanced(root *TreeNode) bool {
	left := height(root.Left)
	right := height(root.Right)

	return max(left, right)-min(left, right) <= 1
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left := height(root.Left)
	right := height(root.Right)

	return max(left, right) + 1
}

func main() {
	tree := &TreeNode{Val: 50}
	tree.insert(20)
	tree.insert(40)
	tree.insert(80)
	tree.insert(79)
	tree.insert(88)
	tree.insert(52)
	tree.insert(25)
	inOrder(tree)
	fmt.Println("")
	fmt.Println(height(tree))
	fmt.Println(isBalanced(tree))
}