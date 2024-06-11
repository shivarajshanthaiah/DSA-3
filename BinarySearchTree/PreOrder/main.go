package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) insert(val int) {
	if t.value < val {
		if t.right == nil {
			t.right = &TreeNode{value: val}
		} else {
			t.right.insert(val)
		}
	} else if t.value > val {
		if t.left == nil {
			t.left = &TreeNode{value: val}
		} else {
			t.left.insert(val)
		}
	}
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value," ")
	preOrder(root.left)
	preOrder(root.right)
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(20)
	tree.insert(49)
	tree.insert(60)
	tree.insert(76)
	tree.insert(87)
	tree.insert(88)
	tree.insert(32)

	preOrder(tree)
	fmt.Println()
}

