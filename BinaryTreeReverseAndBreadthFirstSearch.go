package main

import (
	"log"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Value: val}
}

func main() {
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
	//q := NewQueue()
	//q.PushBack(NewTreeNode(1))
	//q.PushBack(NewTreeNode(2))
	//q.PushBack(NewTreeNode(3))
	//q.Println()
	//q.PopFront()
	//q.PushBack(NewTreeNode(4))
	//q.Println()
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

type Queue struct {
	Front *QueueNode
	Rear  *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Clear() {
	q.Front, q.Rear = nil, nil
}

func (q *Queue) Println() {
	for n := q.Front; nil != n; n = n.Next {
		log.Println(n.Value)
	}
}

func (q *Queue) PopFront() *TreeNode {
	n := q.Front
	q.Front = q.Front.Next
	q.Front.Pre = nil

	n.Next = nil
	n.Pre = nil

	return n.Value
}

func (q *Queue) PushBack(val *TreeNode) {
	if q.Front == q.Rear && q.Front == nil {
		n := NewQueueNode(val)
		q.Front = n
		q.Rear = n
		return
	}

	n := NewQueueNode(val)
	n.Pre = q.Rear
	q.Rear.Next = n
	q.Rear = n
}

type QueueNode struct {
	Value *TreeNode
	Next  *QueueNode
	Pre   *QueueNode
}

func NewQueueNode(v *TreeNode) *QueueNode {
	return &QueueNode{Value: v}
}
