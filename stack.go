package main

import "log"

func Go_Stack() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	log.Println(s.Pop())
	log.Println(s.Pop())
	log.Println(s.Pop())
}

//stack base on array
type Stack struct {
	DataArray []int
	Head      int //index of DataArray
}

const DefaultStackSize = 1024 * 2 //2k
func NewStack() *Stack {
	return &Stack{DataArray: make([]int, DefaultStackSize), Head: -1}
}

func (s *Stack) Pop() int {
	n := s.DataArray[s.Head]
	s.Head -= 1
	return n
}

func (s *Stack) Push(v int) {
	if s.Head+1 >= len(s.DataArray) {
		//morestack
		s.DataArray = append(s.DataArray, make([]int, len(s.DataArray))...)
	}
	s.Head += 1
	s.DataArray[s.Head] = v
}

//stack base on linkedlist
//type Stack struct {
//	Head *StackNode
//}
//
//func NewStack() *Stack {
//	return &Stack{Head: nil}
//}
//
//func (s *Stack) Pop() int {
//	n := s.Head
//	s.Head = s.Head.Next
//	n.Next = nil
//	return n.Value
//}
//
//func (s *Stack) Push(v int) {
//	n := NewStackNode(v)
//	if nil == s.Head {
//		s.Head = n
//		return
//	}
//
//	n.Next = s.Head
//	s.Head = n
//}
//
//type StackNode struct {
//	Value int
//	Next  *StackNode
//}
//
//func NewStackNode(v int) *StackNode {
//	return &StackNode{Value: v}
//}
