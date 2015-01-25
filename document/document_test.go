package document_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/term"
	"testing"
)

func TestDocument(t *testing.T) {
	wantID := 42
	wantTerms := new(term.HashBag)
	wantTerms.Add(term.T("hello"), term.T("world"))

	d := document.NewD(wantID, wantTerms)

	if want, got := wantID, d.ID(); want != got {
		t.Fatalf("want id %v, got %v", want, got)
	}

	if want, got := wantTerms, d.Terms(); want != got {
		t.Fatalf("want terms %v, got %v", want, got)
	}
}
