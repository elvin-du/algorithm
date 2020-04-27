package main

import "log"

func Go_Sort() {
	a := []int{7, 9, 2, 50, 1}
	//BubbleSort(a)
	SelectSort(a)
	log.Println(a)
}
func BubbleSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func SelectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(a)-1; j++ {
			if a[maxIndex] < a[j] {
				maxIndex = j
			}
		}

		a[maxIndex], a[i] = a[i], a[maxIndex]
	}
}
