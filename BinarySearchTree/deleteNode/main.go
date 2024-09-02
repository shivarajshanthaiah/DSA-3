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

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.value, " ")
	inOrder(root.right)
}

func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.value {
		root.left = deleteNode(root.left, val)
	} else if val > root.value {
		root.right = deleteNode(root.right, val)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		successor := findMin(root.right)
		root.value = successor.value
		root.right = deleteNode(root.right, successor.value)
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(40)
	tree.insert(30)
	tree.insert(89)
	tree.insert(8)
	tree.insert(99)
	tree.insert(23)
	tree.insert(96)

	inOrder(tree)
	fmt.Println()

	tree = deleteNode(tree, 89)
	inOrder(tree)
	fmt.Println()
}
