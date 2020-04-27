package main

import "testing"

func TestLinkedList_Append(t *testing.T) {
	l := NewLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Print()
}
