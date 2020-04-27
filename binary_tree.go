package main

import "log"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Value: val}
}

func Go_BinaryTree() {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)

	queue := NewQueue()
	Order(queue, root)
	Reverse(root)

	queue.Clear()
	Order(queue, root)

}

func Order(queue *Queue, root *TreeNode) {
	if nil == root {
		return
	}
	log.Println(root.Value)

	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	r := queue.PopFront()
	Order(queue, r)
}

func Reverse(root *TreeNode) {
	if nil == root {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	Reverse(root.Left)
	Reverse(root.Right)
}
