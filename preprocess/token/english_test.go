package token_test

import (
	"github.com/aybabtme/search/preprocess/token"
	"github.com/aybabtme/search/term"
	"reflect"
	"strings"
	"testing"
)

func TestEnglish(t *testing.T) {
	input := `Convert all documents in collection D to tf-idf weighted vectors, dj, for keyword vocabulary V.
Convert query to a tf-idf-weighted vector q. For each dj in D do
Compute score sj = cosSim(dj, q)
Sort documents by decreasing score. Present top ranked documents to the user.`
	tokens := []string{
		"Convert", "all", "documents", "in", "collection", "D", "to", "tf-idf", "weighted", "vectors", "dj", "for", "keyword", "vocabulary", "V",
		"Convert", "query", "to", "a", "tf-idf-weighted", "vector", "q", "For", "each", "dj", "in", "D", "do",
		"Compute", "score", "sj", "cosSim", "dj", "q",
		"Sort", "documents", "by", "decreasing", "score", "Present", "top", "ranked", "documents", "to", "the", "user",
	}

	wantTerms := new(term.HashBag)
	for _, word := range tokens {
		wantTerms.Add(term.T(word))
	}

	tk := new(token.English)

	gotTerms, err := tk.Tokenize(strings.NewReader(input))
	if err != nil {
		t.Fatalf("tokenizing: %v", err)
	}

	if !reflect.DeepEqual(wantTerms, gotTerms) {
		t.Logf("want=%#v", wantTerms)
		t.Logf(" got=%#v", gotTerms)
		t.Fatalf("mismatch!")
	}
}
