package main

import "log"

func Go_HeapSort() {
	arr := []int{88, 11, 20, 30, 50}
	HeapSort(arr)
	log.Println(arr)
}

func HeapSort(a []int) {
	for l := len(a) - 1; l >= 0; l-- {
		buildMaxHeap(a, l+1)
		a[l], a[0] = a[0], a[l]
	}
}

func buildMaxHeap(a []int, l int) {
	for i := l/2 - 1; i >= 0; i-- {
		heapify(a, i, l)
	}
}

func heapify(a []int, i, l int) {
	left := 2*i + 1
	right := 2*i + 2

	largestIndex := i
	if left < l {
		if a[left] > a[largestIndex] {
			largestIndex = left
		}
	}

	if right < l {
		if a[right] > a[largestIndex] {
			largestIndex = right
		}
	}

	if largestIndex != i {
		a[i], a[largestIndex] = a[largestIndex], a[i]
		heapify(a, largestIndex, l)
	}
}
