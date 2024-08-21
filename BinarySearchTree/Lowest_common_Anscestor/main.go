package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) insert(val int) {
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right.insert(val)
		}
	} else if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Left.insert(val)
		}
	}
}

func lowestCommonAscestor(root, p, q *TreeNode) int {
	if root == nil {
		return -1
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAscestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAscestor(root.Right, p, q)
	}
	return root.Val
}

func main() {
	tree := &TreeNode{Val: 50}
	tree.insert(20)
	tree.insert(40)
	tree.insert(80)
	tree.insert(79)
	tree.insert(88)
	tree.insert(52)
	tree.insert(25)
	tree.insert(38)
	tree.insert(77)
	tree.insert(41)
	tree.insert(64)
	tree.insert(14)
	tree.insert(69)

	p := tree.Left.Left
	q := tree.Left.Right

	fmt.Println(lowestCommonAscestor(tree, p, q))
}
