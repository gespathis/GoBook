package main

import (
	"fmt"
	"strings"
)

func dataSource(uploadStream chan<- string) {

	for _, v := range []string{"fight club", "do", "not talk about", "first rule of fight"} {
		uploadStream <- v
	}
	close(uploadStream)
}

func wordSeperator(uploadStream chan<- string, downloadStream <-chan string) {

	for v := range downloadStream {
		for _, word := range strings.Fields(v) {
			uploadStream <- word
		}
	}
	close(uploadStream)
}

func notSameWord(uploadStream chan<- string, downloadStream <-chan string) {
	previousWord := make(map[string]bool)

	for v := range downloadStream {
		if !previousWord[v] {
			previousWord[v] = true
			uploadStream <- v
		}
	}
	close(uploadStream)
}

func printWords(downloadStream <-chan string) {

	for {
		word, ok := <-downloadStream

		if !ok {
			return
		}
		fmt.Printf("the word is : %v \n", word)
	}
}

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	c2 := make(chan string)

	go dataSource(c0)
	go wordSeperator(c1, c0)
	go notSameWord(c2, c1)
	printWords(c2)
}
