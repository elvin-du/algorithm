```
package main

import (
	"fmt"
	"github.com/go-errors/errors"
	"log"
	"math/rand"
	"time"
)

func main() {
	root := InitBTree(5)
	InOrderBTree(root)
	val,err := Search(55,root)
	if nil != err{
		log.Println(err)
		return
	}

	log.Println("Found value:",val)
}

func Search(key int32, root *Node) (interface{}, error) {
	for {
		if nil == root {
			break
		} else if key == root.Key {
			return root.Value, nil
		} else if key > root.Key {
			root = root.Right
			continue
		} else {
			root = root.Left
		}
	}

	return "", errors.New("Not Found")
}

func randInt(size int) []int32 {
	rand.Seed(time.Now().UnixNano())

	data := make([]int32, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, rand.Int31n(100))
	}
	return data
}

func InOrderBTree(root *Node) {
	if nil == root {
		return
	}
	InOrderBTree(root.Left)
	ProcessNode(root)
	InOrderBTree(root.Right)
}

func ProcessNode(n *Node) {
	log.Println("key:", n.Key, "value:", n.Value)
}

func InitBTree(size int) *Node {
	data := randInt(size)
	log.Println("data:", data)

	root := NewNode(data[0], fmt.Sprintf("0x%X", data[0]))
	for i := 1; i < size; i++ {
		n := NewNode(data[i], fmt.Sprintf("0x%X", data[i]))

		InsertNode(root, n)
	}

	return root
}

func InsertNode(parent *Node, newNode *Node) {
	for {
		if newNode.Key >= parent.Key {
			if parent.Right == nil {
				parent.Right = newNode
				return
			} else {
				parent = parent.Right
				//InsertNode(parent.Right, newNode)
			}
		} else {
			if parent.Left == nil {
				parent.Left = newNode
				return
			} else {
				parent = parent.Left
				//InsertNode(parent.Left, newNode)
			}
		}
	}
}

func NewNode(key int32, value interface{}) *Node {
	return &Node{Key: key, Value: value}
}

type Node struct {
	Right *Node
	Left  *Node
	Key   int32
	Value interface{}
}


```
