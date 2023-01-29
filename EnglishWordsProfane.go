package enwords

import "strings"

func EnglishWordsProfane() []string {
	wordsProfane = strings.TrimSpace(wordsProfane)
	wordProfaneList := strings.Split(wordsProfane, "\n")
	return wordProfaneList
}
