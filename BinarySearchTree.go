package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
	"math/rand"
	"time"
)

var db *leveldb.DB = nil

func main() {
	var err error = nil
	db, err = leveldb.OpenFile("./db", nil)
	if nil != err {
		panic(err)
	}

	root := InitBTree(5)
	InOrderBTree(root)
	val, err := Search(55, root)
	if nil != err {
		log.Println(err)
	} else {
		log.Println("Found value:", val)
	}

	CommitDB(root)
	IteratorDBTree(root.hash)
}

//按照从小到大的顺序读取出来
func IteratorDBTree(root string) {
	if "" == root {
		return
	}

	value, err := db.Get([]byte(root), nil)
	if nil != err && leveldb.ErrNotFound != err {
		panic(err)
	} else if leveldb.ErrNotFound == err {
		log.Printf("not found in db,key:%s\n", root)
		return
	}

	log.Printf("raw db data:%+v,key:%s", string(value), string(root))

	node := struct {
		Left  string
		Right string
		Key   int32
		Value interface{}
	}{}

	err = json.Unmarshal(value, &node)
	if nil != err {
		panic(err)
	}

	IteratorDBTree(node.Left)
	log.Printf("Unmarshal node:%+v\n", node)

	IteratorDBTree(node.Right)
}

func CommitDB(root *Node) {
	if nil == root {
		return
	}

	CommitDB(root.Left)
	CommitDB(root.Right)
	_, dbVal := root.Sha3()
	log.Printf("commitDB key:%s,value:%s\n", root.hash, dbVal)
	err := db.Put([]byte(root.hash), dbVal, &opt.WriteOptions{Sync: true})
	if nil != err {
		panic(err)
	}
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
	root.hash, _ = root.Sha3()
	for i := 1; i < size; i++ {
		n := NewNode(data[i], fmt.Sprintf("0x%X", data[i]))
		n.hash, _ = n.Sha3()
		InsertNode(root, n)
	}

	return root
}

func InsertNode(parent *Node, newNode *Node) {
	for {
		if newNode.Key >= parent.Key {
			if parent.Right == nil {
				parent.Right = newNode
				//update hash
				parent.hash, _ = parent.Sha3()
				return
			} else {
				parent = parent.Right
				//InsertNode(parent.Right, newNode)
			}
		} else {
			if parent.Left == nil {
				parent.Left = newNode
				parent.hash, _ = parent.Sha3()
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
	hash  string
	Key   int32
	Value interface{}
}

func sha3(data interface{}) (key string, value []byte) {
	bin, err := json.Marshal(data)
	if nil != err {
		panic(err)
	}
	k := sha256.Sum256(bin)
	return hex.EncodeToString(k[:]), bin
}

func (n *Node) Unmarshal(bin []byte) {
	tmp := struct {
		Left  string
		Right string
		Key   int32
		Value interface{}
	}{}

	err := json.Unmarshal(bin, &tmp)
	if nil != err {
		panic(err)
	}

	//n.hash =
}
func (n *Node) Sha3() (key string, value []byte) {
	left, right := "", ""
	if n.Right != nil {
		right = n.Right.hash
	}

	if n.Left != nil {
		left = n.Left.hash
	}
	tmp := struct {
		Left  string
		Right string
		Key   int32
		Value interface{}
	}{left, right, n.Key, n.Value}

	return sha3(tmp)
}
