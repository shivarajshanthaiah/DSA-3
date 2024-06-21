package main

import (
	"fmt"
	"math"
)

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

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.value <= min || node.value >= max {
		return false
	}

	return isValid(node.left, min, node.value) && isValid(node.right, node.value, max)
}


func main() {
	tree := &TreeNode{value: 50}
	tree.insert(40)
	tree.insert(40)
	tree.insert(22)
	tree.insert(90)
	tree.insert(-1)
	tree.right = &TreeNode{value: 7}

	fmt.Println(isValidBST(tree))
}
