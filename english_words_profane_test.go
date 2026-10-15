package enwords

import (
	"testing"
)

func TestEnglishWordsProfane(t *testing.T) {
	wordsProfane := EnglishWordsProfane()

	if len(wordsProfane) < 2 {
		t.Error("There must be at least 2 words in the list but found: ", len(wordsProfane))
	}

	if wordsProfane[0] != "2 girls 1 cup" {
		t.Error("First word MUST BE '2', but found: ", "'"+wordsProfane[0]+"'")
	}

	if wordsProfane[len(wordsProfane)-1] != "🖕" {
		t.Error("Last word MUST BE '🖕', but found: ", "'"+wordsProfane[len(wordsProfane)-1]+"'")
	}
}
