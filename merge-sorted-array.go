package main

import "log"

var (
	a = [4]int{4, 8, 18, 66}
	b = [8]int{7, 10, 12, 77}
)

func Go_MergeSortedArray() {
	MergeSortedArray(a[:], b[:], 4, 4)
	log.Println(b)
}

func MergeSortedArray(a, b []int, alen, blen int) {
	newLen := len(b) - 1

	ai := alen - 1
	bi := blen - 1
	for ; ai >= 0; ai-- {
		for ; bi >= 0; bi-- {
			if a[ai] > b[bi] {
				b[newLen] = a[ai]
				newLen -= 1
				break
			} else {
				b[newLen] = b[bi]
				newLen -= 1
			}
		}

		if bi < 0 {
			break
		}
	}

	if bi < 0 {
		copy(b[0:ai+1], a[0:ai+1])
	}
}
