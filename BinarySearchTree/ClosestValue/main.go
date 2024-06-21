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

func findClosestVal(node *TreeNode, target int) int {
	closest := node.value
	temp := node

	for temp != nil {
		if temp.value == target {
			return target
		}

		if abs(target-temp.value) < abs(target-closest) {
			closest = temp.value
		}
		if target < temp.value {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	tree := &TreeNode{value: 50}
	tree.insert(57)
	tree.insert(67)
	tree.insert(32)
	tree.insert(22)
	tree.insert(45)
	tree.insert(78)

	fmt.Println(findClosestVal(tree, 70))
}
