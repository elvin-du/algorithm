package main

import (
	"algorithm/sort"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	sort.LoadRawData()

	var sorter sort.Sorter

	sorter = sort.NewSelection()

	err := sorter.Sort()
	if nil != err {
		log.Fatalln(err)
	}

	if sorter.IsSorted() {
		sorter.Show(10)
	}
}
