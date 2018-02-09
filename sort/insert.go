package sort

import (
	"log"
	"time"
)

type Insert struct{}

func NewInsert() *Insert {
	return &Insert{}
}

func (i *Insert) Sort() error {
	log.Println("Insert sort start...")
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
	log.Println("Insert sort done...")
	log.Printf("cost time: %v", time.Now().Sub(start))
	return nil
}

func (i *Insert) IsSorted() bool {
	for i := 0; i < gDataLength-1; i++ {
		if gRawData[i] >= gRawData[i+1] {
			return false
		}
	}
	return true
}

func (i *Insert) Show(limit int) {
	log.Printf("order by from small to large,only show first %d", limit)
	for i := 0; i < limit; i++ {
		log.Println(gRawData[i])
	}
}
