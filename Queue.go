package main

import "log"

func Go_Queue() {
	q := NewQueue()
	q.PushBack(NewTreeNode(1))
	q.PushBack(NewTreeNode(2))
	q.PushBack(NewTreeNode(3))
	q.Println()
	q.PopFront()
	q.PushBack(NewTreeNode(4))
	q.Println()
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
