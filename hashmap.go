package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

func Go_Hashmap() {
	hmp := NewHashmap()
	hmp.Add("123", "456")
	hmp.Add("abc", "eeefff")
	hmp.Add(55, 66)
	hmp.Add(77, 88)

	log.Println(hmp.Get("123"))
	log.Println(hmp.Get("abc"))
	log.Println(hmp.Get(55))
	log.Println(hmp.Get(77))

	removede := hmp.Delete("123")
	log.Println("removed element", removede)
	log.Println(hmp.Get("123"))

}

type Hashmap struct {
	Data []interface{}
}

type bucket struct {
	Key   interface{}
	Value interface{}
	Next  *bucket
}

func NewBucket(key, value interface{}) *bucket {
	return &bucket{Key: key, Value: value}
}

const defaultHashmapSize = 1024

func NewHashmap() *Hashmap {
	return &Hashmap{Data: make([]interface{}, defaultHashmapSize)}
}

func (h *Hashmap) Get(key interface{}) interface{} {
	data := h.Data[h.hash(key)]
	if nil == data {
		return nil
	}

	for b := data.(*bucket); nil != b; b = b.Next {
		if b.Key == key {
			return b.Value
		}
	}

	return nil
}

func (h *Hashmap) Delete(key interface{}) interface{} {
	hashindex := h.hash(key)
	data := h.Data[hashindex]
	if nil == data {
		return nil
	}

	b := data.(*bucket)
	if b.Key == key {
		h.Data[hashindex] = nil
		return b.Value
	}

	vhead := NewBucket(nil, nil)
	vhead.Next = b
	pre, cur := vhead, vhead.Next
	for ; cur != nil; pre, cur = pre.Next, cur.Next {
		if cur.Key == key {
			pre.Next = pre.Next.Next
			cur.Next = nil
			return cur.Value
		}
	}

	return nil
}

func (h *Hashmap) hash(key interface{}) int {
	switch v := key.(type) {
	case string:
		key = []byte(v)
	case int:
		key = uint64(v)
	case uint:
		key = uint64(v)
	}

	w := bytes.NewBuffer(nil)
	err := binary.Write(w, binary.LittleEndian, key)
	if nil != err {
		panic(err)
	}

	sha := sha256.New()
	sha.Write(w.Bytes())
	binSha := sha.Sum(nil)

	//64bit
	int64Sha, n := binary.Uvarint(binSha[:8])
	if n <= 0 {
		panic("convert []byte to uint64 failed")
	}

	return int(int64Sha % defaultHashmapSize)
}

func (h *Hashmap) doUpdate(firstb, newb *bucket) {
	vHead := NewBucket(nil, nil)
	vHead.Next = firstb
	pre, cur := vHead, vHead.Next

	for ; cur != nil; pre, cur = pre.Next, cur.Next {
		if cur.Key == newb.Key {
			//update
			cur.Value = newb.Value
			return
		}
	}

	//add
	pre.Next = newb
}

func (h *Hashmap) Add(key interface{}, val interface{}) {
	b := NewBucket(key, val)
	hash := h.hash(key)
	firstBucket := h.Data[hash]
	if nil == firstBucket {
		h.Data[hash] = b
	} else {
		h.doUpdate(firstBucket.(*bucket), b)
	}
}
