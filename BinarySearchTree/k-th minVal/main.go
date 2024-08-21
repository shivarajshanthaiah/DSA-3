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

func kthMinVal(root *TreeNode, n int, k *int) {
	if root == nil {
		return
	}

	kthMinVal(root.Left, n, k)
	*k++
	if *k == n {
		fmt.Println(root.Val)
	}

	kthMinVal(root.Right, n, k)
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
	k := 0
	n := 5
	kthMinVal(tree, n, &k)
}
