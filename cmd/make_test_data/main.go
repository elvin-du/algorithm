package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var filePath = "../../data.txt"

func main() {
	f, err := os.Create(filePath)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	length := 10000 * 100
	for i := 0; i < length; i++ {
		_, err = w.WriteString(fmt.Sprintf("%d:", rand.Uint64()))
		if nil != err {
			log.Fatalln(err)
		}
	}
	log.Println("make success")
}
