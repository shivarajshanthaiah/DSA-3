package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (tree *TreeNode) insert(val int) {
	if tree.value < val {
		if tree.right == nil {
			tree.right = &TreeNode{value: val}
		} else {
			tree.right.insert(val)
		}
	} else if tree.value > val {
		if tree.left == nil {
			tree.left = &TreeNode{value: val}
		} else {
			tree.left.insert(val)
		}
	}
}

func (root *TreeNode) contains(key int) bool {
	if root == nil {
		return false
	}

	if root.value < key {
		return root.right.contains(key)
	} else if root.value > key {
		return root.left.contains(key)
	}

	return true
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(57)
	tree.insert(67)
	tree.insert(32)
	tree.insert(22)
	tree.insert(45)
	tree.insert(78)

	fmt.Println(tree.contains(7))
}
