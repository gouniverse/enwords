package enwords

import "strings"

func EnglishWords() []string {
	words = strings.TrimSpace(words)
	wordList := strings.Split(words, "\n")
	return wordList
}
