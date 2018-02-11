package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var filePath = "../sort/data.txt"

func main() {
	f, err := os.Create(filePath)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	//	length := 10000 * 100
	length := 10000 * 10
	for i := 0; i < length; i++ {
		_, err = w.WriteString(fmt.Sprintf("%d\n", rand.Uint64()))
		if nil != err {
			log.Fatalln(err)
		}
	}
	w.Flush()
	log.Println("make success")
}
