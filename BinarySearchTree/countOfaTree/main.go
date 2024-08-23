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

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := count(root.Left)
	right := count(root.Right)
	return left + right + 1
}

func kthLevelVals(root *TreeNode, level, k int) {
	if root == nil {
		return
	}

	if level == k {
		fmt.Print(" ", root.Val)
		return
	}

	kthLevelVals(root.Left, level+1, k)
	kthLevelVals(root.Right, level+1, k)
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
	tree.insert(63)
	inOrder(tree)
	fmt.Println()
	fmt.Println(count(tree))

	k := 4
	kthLevelVals(tree, 1, k)
	fmt.Println("")
}
