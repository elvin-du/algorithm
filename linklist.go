package main

import "log"

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) ReverseTwo() {

}

func (l *LinkedList) Print() {
	for n := l.Head; nil != n; n = n.Next {
		log.Println(n.Value)
	}
}

func (l *LinkedList) Append(v int) {
	vhead := NewListNode(0)
	vhead.Next = l.Head
	for ; nil != vhead.Next; vhead = vhead.Next {
	}

	vhead.Next = NewListNode(v)
}

func NewListNode(v int) *ListNode {
	return &ListNode{Value: v}
}

type ListNode struct {
	Value int
	Next  *ListNode
}
