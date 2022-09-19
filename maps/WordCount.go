package main

import (
	"strings"
)

func getWordCounts(sentence string) map[string]int {

	wordCounts := make(map[string]int)
	wordsSlice := strings.Split(sentence, " ")
	for _, word := range wordsSlice {
		_, ok := wordCounts[word]
		if ok == true {
			wordCounts[word] += 1
		} else {
			wordCounts[word] = 1
		}
	}
	return wordCounts
}
