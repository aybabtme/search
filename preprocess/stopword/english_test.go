package stopword

import (
	"github.com/aybabtme/search/term"
	"testing"
)

func TestEnglish(t *testing.T) {
	notStopwords := []string{
		"hello", "obiwan", "kenobi",
	}

	stop := new(English)
	for _, word := range englishStopwords {
		if !stop.IsStopWord(term.T(word)) {
			t.Fatalf("%q should be a stop word", word)
		}
	}

	for _, word := range notStopwords {
		if stop.IsStopWord(term.T(word)) {
			t.Fatalf("%q should not be a stop word", word)
		}
	}
}
