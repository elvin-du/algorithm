package main

import "log"

func Go_QuickSort() {
	a := []int{88, 1, 6, 7, 20, 33}
	QuickSort(a, 0, len(a)-1)
	log.Println(a)
}

func QuickSort(a []int, s, e int) {
	if s >= e {
		return
	}

	left, right := s, e
	pivot := a[s]
	for ; s < e; {
		for ; e > s; e-- {
			if a[e] < pivot {
				a[s] = a[e]
				break
			}
		}

		for ; s < e; s++ {
			if a[s] > pivot {
				a[e] = a[s]
				break
			}
		}
	}
	a[s] = pivot

	QuickSort(a, left, s-1)
	QuickSort(a, s+1, right)
}
