package main

import (
	"bytes"
	"log"
	"unicode/utf8"
)

var urlcharactor = "we are happy"
var oldCharacter = ' '

func Go_UrlCharacterReplace() {
	log.Println(UrlCharacterReplace(urlcharactor))
}

func UrlCharacterReplace(str string) string {
	oldLen := utf8.RuneCountInString(str)
	newLen := oldLen
	for _, r := range str {
		if oldCharacter == r {
			newLen += 2
		}
	}

	oldStr := bytes.Runes([]byte(str))
	newStr := make([]rune, newLen)
	for i, j := oldLen-1, newLen-1; i >= 0; i-- {
		if oldCharacter == oldStr[i] {
			newStr[j] = '0'
			newStr[j-1] = '2'
			newStr[j-2] = '%'
			j -= 3
			continue
		}

		newStr[j] = oldStr[i]
		j -= 1
	}

	return string(newStr)
}
