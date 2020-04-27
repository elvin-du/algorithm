package main

import "log"

func Go_LinkedList() {
	l := NewLinkedList()
	l.Append(1)
	l.Append(4)
	l.Append(5)
	l.Print()
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) ReverseTwo() {
	pre, cur, next := (*ListNode)(nil), l.Head, l.Head.Next
	for ; cur != nil; {
		cur1 := next.Next
		next.Next = cur
		cur.Next = pre

		pre = next
		cur = cur1
		next = cur.Next
	}
}

func (l *LinkedList) Print() {
	for n := l.Head; nil != n; n = n.Next {
		log.Println(n.Value)
	}
}

func (l *LinkedList) Append(v int) {
	vhead := NewListNode(0)
	vhead.Next = l.Head
	node := vhead
	for ; nil != node.Next; node = node.Next {
	}

	node.Next = NewListNode(v)
	l.Head = vhead.Next
}

func NewListNode(v int) *ListNode {
	return &ListNode{Value: v}
}

type ListNode struct {
	Value int
	Next  *ListNode
}
