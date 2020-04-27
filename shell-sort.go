package main

import "log"

func Go_ShellSort() {
	arr := []int{88, 77, 1, 99, 100, 200, 8}
	ShellSort(arr, len(arr)/2)
	log.Println(arr)
}
func ShellSort(arr []int, k int) {
	if k < 0 {
		return
	}
	InsertSort(arr, k-1)
}
func InsertSort(arr []int, k int) {
	//k := len(arr) / 2
	if k < 0 {
		return
	}

	for i := 1; i < len(arr); i += k {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0; j -= k {
			if tmp >= arr[j] {
				break
			} else if tmp < arr[j] {
				arr[j+1] = arr[j]
			}
		}

		arr[j+1] = tmp
	}
}
