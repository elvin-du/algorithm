package main
//
//import (
//	"fmt"
//	"log"
//)
//
//func init() {
//	log.SetFlags(log.Lshortfile)
//}
//func main() {
//	list := NewLinkedList()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Print()
//
//	log.Println(list.Find(3).Next)
//
//	list.Delete(3)
//	list.Print()
//}
//
//type LinkedList struct {
//	Head *ListNode
//	Tail *ListNode
//}
//
//func NewLinkedList() *LinkedList {
//	tmp := NewListNode(0)
//	return &LinkedList{Head: tmp, Tail: tmp}
//}
//
//func (l *LinkedList) Append(d interface{}) {
//	n := NewListNode(d)
//	n.Pre = l.Tail
//	l.Tail.Next = n
//	l.Tail = n
//	l.Head.Data = l.Head.Data.(int) + 1
//}
//
//func (l *LinkedList) Prepend(d interface{}) {
//	n := NewListNode(d)
//	n.Pre = l.Head
//	n.Next = l.Head.Next
//	l.Head.Next = n
//
//	l.Head.Data = l.Head.Data.(int) + 1
//}
//
//func (l *LinkedList) Find(d interface{}) *ListNode {
//	for tmp := l.Head.Next; nil != tmp; tmp = tmp.Next {
//		if tmp.Data == d {
//			return tmp
//		}
//	}
//
//	return nil
//}
//func (l *LinkedList) Delete(d interface{}) {
//	ret := l.Find(d)
//	ret.Pre.Next= ret.Next
//}
//
//func (l *LinkedList) Print() {
//	log.Println("total elements:", l.Head.Data)
//	for iter := l.Head.Next; iter != nil; iter = iter.Next {
//		log.Println(iter.Data)
//	}
//}
//
//type ListNode struct {
//	Data interface{}
//	Next *ListNode
//	Pre  *ListNode
//}
//
//func NewListNode(d interface{}) *ListNode {
//	return &ListNode{Data: d}
//}
//
//func (n *ListNode) String() string {
//	return fmt.Sprintf("data:%+v", n.Data)
//}
