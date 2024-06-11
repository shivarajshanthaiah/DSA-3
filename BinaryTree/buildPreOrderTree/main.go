package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func buildTree(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := 0
	return buildTreeHelper(preorder, &index, len(preorder))
}

func buildTreeHelper(preorder []int, index *int, n int) *TreeNode {
	if *index >= n {
		return nil
	}

	if preorder[*index] == -1 {
		*index++
		return nil
	}

	root := &TreeNode{val: preorder[*index]}
	*index++

	root.left = buildTreeHelper(preorder, index, n)
	root.right = buildTreeHelper(preorder, index, n)

	return root
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.val)
		helper(node.left)
		helper(node.right)
	}
	helper(root)
	return result
}

func main() {
	preorder := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, -1}
	root := buildTree(preorder)

	res := preorderTraversal(root)
	fmt.Println(res)
}
