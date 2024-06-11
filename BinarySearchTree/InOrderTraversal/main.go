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

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(" ",root.value)
	inOrder(root.right)
}

func main() {
	tree := &TreeNode{value: 100}
	tree.insert(30)
	tree.insert(59)
	tree.insert(67)
	tree.insert(201)
	tree.insert(516)
	tree.insert(782)
	tree.insert(10)
	tree.insert(88)

	inOrder(tree)

}
