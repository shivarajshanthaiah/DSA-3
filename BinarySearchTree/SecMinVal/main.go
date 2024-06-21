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
	fmt.Print(" ", node.value)
	inOrder(node.right)
}

func findSecMin(root *TreeNode) int {
	min := math.MaxInt64
	secMin := math.MaxInt64
	tMin := math.MaxInt64

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.value < min {
			tMin = secMin
			secMin = min
			min = node.value
		} else if node.value < secMin && node.value > min {
			tMin = secMin
			secMin = node.value
		} else if node.value < tMin && node.value > secMin && node.value > min {
			tMin = node.value
		}
		traverse(node.left)
		traverse(node.right)
	}
	traverse(root)
	return tMin
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
	fmt.Println(findSecMin(tree))
}
