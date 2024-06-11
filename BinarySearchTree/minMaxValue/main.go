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

func findMin(node *TreeNode) *TreeNode {
	temp := node

	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func findMax(node *TreeNode) *TreeNode {
	temp := node
	for temp.right != nil {
		temp = temp.right
	}
	return temp
}

func main() {
	tree := &TreeNode{value: 50}

	tree.insert(59)
	tree.insert(30)
	tree.insert(28)
	tree.insert(88)
	tree.insert(76)
	tree.insert(10)

	min := findMin(tree)
	if min != nil {
		fmt.Println("Minimum value is :", min.value)
	} else {
		fmt.Println("Tree is empty")
	}

	max := findMax(tree)
	fmt.Println("Max value is :", max.value)
}
