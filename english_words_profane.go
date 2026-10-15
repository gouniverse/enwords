package enwords

import "strings"

func EnglishWordsProfane() []string {
	words_profane = strings.TrimSpace(words_profane)
	wordProfaneList := strings.Split(words_profane, "\n")
	return wordProfaneList
}
