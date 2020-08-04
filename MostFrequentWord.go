package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string = "I like chocolate, vanilla, and lemon and orange ice cream."
	fmt.Println("The most frequent word in this sentence is :")
	fmt.Println(max(theWordCount(sentence)))
}

func theWordCount(sentence string) map[string]int {
	words := strings.Fields(sentence)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

func max(sentence map[string]int) (string, int) {
	maxCount := 0
	maxWord := " "
	for word, count := range sentence {
		if count > maxCount {
			maxCount = count
			maxWord = word

		}
	}
	return maxWord, maxCount
}
