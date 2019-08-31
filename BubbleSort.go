package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	data := RandInt32(5)
	log.Println("raw data:", data)

	BubbleSort(&data)
	log.Println("sorted data:", data)
}

func RandInt32(size int) []int32 {
	data := make([]int32, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, rand.Int31n(int32(100)))
	}

	return data
}

func BubbleSort(data *[]int32) {
	size := len(*data)
	for i := size - 1; i >= 0; i-- {
		maxIndex := FindMax((*data)[0 : i+1])
		(*data)[maxIndex], (*data)[i] = (*data)[i], (*data)[maxIndex]
	}
}

func FindMax(data []int32) int {
	maxIndex := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}
