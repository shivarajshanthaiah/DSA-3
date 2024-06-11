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

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left := height(root.left)
	right := height(root.right)

	return max(left, right) + 1
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(70)
	tree.insert(20)
	tree.insert(25)
	tree.insert(89)
	tree.insert(55)
	tree.insert(10)
	tree.insert(42)
	tree.insert(75)
	tree.insert(99)
	tree.insert(60)
	tree.insert(92)

	fmt.Println(height(tree))
}


