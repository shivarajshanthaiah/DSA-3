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

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(20)
	tree.insert(40)
	tree.insert(80)
	tree.insert(79)
	tree.insert(88)
	tree.insert(52)
	tree.insert(25)

	bfs(tree)
	fmt.Println()

}
