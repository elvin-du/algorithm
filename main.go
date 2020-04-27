//package main
//
//import "log"
//
//func init() {
//	//log.setflags(log.lshortfile)
//}
//func main() {
//	//go_hashmap()
//	//go_stack()
//	//
//	//cache := make([]*frame,3)
//	//cache[0].ip = 8
//	//log.println(cache[0])
//
//	//go_urlcharacterreplace()
//	//go_mergesortedarray()
//	//go_linkedlist()
//	//Go_Sort()
//	//Go_ShellSort()
//	//GoMergeSort()
//	//Go_HeapSort()
//	//Go_QuickSort()
//	s := "abcd杜满想"
//	r := []rune(s)
//	log.Println(string(r[5]))
//	//for _,v := range r{
//	//	log.Printf("%s",string( v))
//	//}
//
//}
//
//type frame struct {
//	ip   int
//	name string
//}

package main

import (
	"fmt"
	"reflect"
)
type A struct {
	Name string
}
func main() {
	var v interface{} = &A{}
	var val interface{} = v
	fmt.Println(reflect.TypeOf(val).Kind())
	val = 50
	fmt.Println(reflect.TypeOf(val))
}

func IsNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}

	return false
}
