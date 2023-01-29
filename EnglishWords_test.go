package enwords

import (
	"testing"
)

func TestEnglishWords(t *testing.T) {
	words := EnglishWords()

	if len(words) < 2 {
		t.Error("There must be at least 2 words in the list but found: ", len(words))
	}

	if words[0] != "2" {
		t.Error("First word MUST BE '2', but found: ", "'"+words[0]+"'")
	}

	if words[len(words)-1] != "ZZZ" {
		t.Error("Last word MUST BE 'ZZZ', but found: ", "'"+words[len(words)-1]+"'")
	}
}
