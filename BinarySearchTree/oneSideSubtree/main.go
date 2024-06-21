package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) insert(val int) {
	if node.value < val {
		if node.right == nil {
			node.right = &TreeNode{value: val}
		} else {
			node.right.insert(val)
		}
	} else if node.value > val {
		if node.left == nil {
			node.left = &TreeNode{value: val}
		} else {
			node.left.insert(val)
		}
	}
}
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Print(" ",node.value)
	inOrder(node.right)
}

func rightSubTree(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node
	for temp != nil {
		fmt.Print(" ", temp.value)
		temp = temp.right
	}
}
func main() {
	tree := &TreeNode{value: 50}
	tree.insert(60)
	tree.insert(65)
	tree.insert(23)
	tree.insert(78)
	tree.insert(99)
	tree.insert(76)
	tree.insert(12)
	tree.insert(43)
	inOrder(tree)
	fmt.Println()
	rightSubTree(tree)
}
