package preprocess_test

import (
	"github.com/aybabtme/search/preprocess"
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
	terms := []string{
		"vector",
		"vector",
		"decreas",
		"collect",
		"to",
		"document",
		"q",
		"tf-idf",
		"score",
		"sj",
		"v",
		"sort",
		"cossim",
		"vocabulari",
		"weight",
		"tf-idf-weight",
		"keyword",
		"user",
		"dj",
		"convert",
		"present",
		"the",
		"d",
		"rank",
		"queri",
		"comput",
		"top",
	}

	wantTerms := new(term.HashBag)
	for _, word := range terms {
		wantTerms.Add(term.T(word))
	}

	proc := new(preprocess.English)

	gotTerms, err := proc.Process(strings.NewReader(input))
	if err != nil {
		t.Fatalf("preprocessing: %v", err)
	}

	if !reflect.DeepEqual(wantTerms, gotTerms) {
		t.Logf("want=%#v", wantTerms)
		t.Logf(" got=%#v", gotTerms)
		t.Fatalf("mismatch!")
	}
}
