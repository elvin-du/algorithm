package main

import "log"

func GoMergeSort() {
	arr := []int{88, 1, 9, 3, 10, 7}
	b := mergeSort(arr)
	log.Println(b)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	a := arr[:mid]
	b := arr[mid:]

	return merge(mergeSort(a), mergeSort(b))
}

func merge(a, b []int) []int {
	c := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for ; i < len(a) && j < len(b); {
		if a[i] < b[j] {
			c = append(c, a[i])
			i += 1
		} else {
			c = append(c, b[j])
			j += 1
		}
	}

	if i >= len(a) {
		c = append(c, b[j:]...)
	}

	if j >= len(b) {
		c = append(c, a[i:]...)
	}

	return c
}
