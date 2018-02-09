package sort

import (
	"log"
	"time"
)

type Selection struct{}

func NewSelection() *Selection {
	return &Selection{}
}

func (s *Selection) Sort() error {
	log.Println("selection sort start...")
	start := time.Now()

	for i := 0; i < gDataLength; i++ {
		flag := i
		for j := i + 1; j < gDataLength; j++ {
			if gRawData[flag] > gRawData[j] {
				flag = j
			}
		}

		if i != flag {
			exchange(i, flag)
		}
	}
	log.Println("selection sort done...")
	log.Printf("cost time: %v", time.Now().Sub(start))
	return nil
}

func (s *Selection) IsSorted() bool {
	for i := 0; i < gDataLength-1; i++ {
		if gRawData[i] > gRawData[i+1] {
			return false
		}
	}
	return true
}

func (s *Selection) Show(limit int) {
	log.Printf("order by from small to large,only show first %d", limit)
	for i := 0; i < limit; i++ {
		log.Println(gRawData[i])
	}
}
