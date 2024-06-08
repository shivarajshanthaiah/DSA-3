package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}

	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

func inTra(root *TreeNode) {
	if root != nil {
		inTra(root.left)
		fmt.Printf("%d", root.val)
		inTra(root.right)
	}
}

func main() {
	var root *TreeNode

	root = insert(root, 3)
	root = insert(root, 6)
	root = insert(root, 9)
	root = insert(root, 5)
	root = insert(root, 7)
	root = insert(root, 2)

	inTra(root)
	fmt.Println()
}
