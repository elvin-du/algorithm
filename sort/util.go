package sort

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	gRawData    = []uint64{}
	gDataLength int
	filePath    = "./data.txt"
)

func LoadRawData() {
	f, err := os.Open(filePath)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if io.EOF == err {
			break
		} else if nil != err {
			log.Fatalln(err)
		}
		num, err := strconv.ParseUint((string(line[:len(line)-1])), 10, 64)
		if nil != err {
			log.Fatalln(err)
		}
		gRawData = append(gRawData, num)
	}
	gDataLength = len(gRawData)
}

func exchange(i, j int) {
	baz := gRawData[i]
	gRawData[i] = gRawData[j]
	gRawData[j] = baz
}
