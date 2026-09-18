package main

import (
	"fmt"
	"sort"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {
	words := make([]string, 0, len(wordMap))
	for word := range wordMap {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if wordMap[words[i]] == wordMap[words[j]] {
			return words[i] < words[j]
		}
		return wordMap[words[i]] > wordMap[words[j]]
	})

	if n > len(words) {
		n = len(words)
	}

	return words[:n]
}

func AnalyzeText(text string) {
	WordsAndCount := make(map[string]int)
	fullCount := 0

	for _, word := range strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '.' || r == ',' || r == '!' || r == '?'
	}) {
		WordsAndCount[strings.ToLower(word)] += 1
		fullCount += 1
	}
	topWords := getTopWords(WordsAndCount, 5)

	fmt.Printf("Количество слов: %d\n", fullCount)
	fmt.Printf("Количество уникальных слов: %d\n", len(WordsAndCount))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], WordsAndCount[topWords[0]])
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	for _, v := range topWords {
		fmt.Printf("\"%s\": %d раз\n", v, WordsAndCount[v])
	}

}
