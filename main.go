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
	fmt.Print(" ", root.value)
	inOrder(root.right)
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left := height(root.left)
	right := height(root.right)
	return max(left, right) + 1
}

func (root *TreeNode) contains(key int) bool {
	if root == nil {
		return false
	}
	if root.value < key {
		return root.right.contains(key)
	} else if root.value > key {
		return root.left.contains(key)
	}
	return true
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
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
		successor := min(root.right)
		root.value = successor.value
		root.right = deleteNode(root.right, successor.value)
	}
	return root
}

func min(root *TreeNode) *TreeNode {
	temp := root
	for temp.left != nil {
		temp = temp.left
	}
	return temp

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
	deleteNode(tree,20)
	inOrder(tree)
	fmt.Println(height(tree))
	fmt.Println(tree.contains(20))
	bfs(tree)
}
