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

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.value," ")
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(30)
	tree.insert(40)
	tree.insert(60)
	tree.insert(66)
	tree.insert(88)
	tree.insert(22)
	tree.insert(55)
	postOrder(tree)
}
