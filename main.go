package main

import (
	"algorithm/sort"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var sorter sort.Sorter

	sorter = sort.NewSelection()

	err := sorter.ReadRawData()
	if nil != err {
		log.Fatalln(err)
	}

	err = sorter.Sort()
	if nil != err {
		log.Fatalln(err)
	}

	if sorter.IsSorted() {
		sorter.Show(10)
	}
}
