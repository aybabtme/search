package index_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/index"
	"github.com/aybabtme/search/term"
	"testing"
)

func TestHashIndex(t *testing.T) {
	testIndexImpl(t, new(index.HashIndex))
}

func testIndexImpl(t testing.TB, idx index.Idx) {
	check := func(numDocs, numTerms int) {
		if want, got := numDocs, idx.TotalNumDocs(); want != got {
			t.Fatalf("want numdocs %v, got %v", want, got)
		}
		if want, got := numTerms, idx.TotalNumTerms(); want != got {
			t.Fatalf("want numterms %v, got %v", want, got)
		}
	}

	// check that the nil value works
	check(0, 0)
	nosuchterm := term.T("nosuchterm")
	if want, got := 0, idx.Get(nosuchterm).NumDocs(); want != got {
		t.Fatalf("want numdocs %v, got %v", want, got)
	}
	check(0, 0)

	A, B, C := term.T("A"), term.T("B"), term.T("C")

	terms := new(term.HashBag).Add(A, B, C)
	wantDoc := document.NewD(0, terms)

	idx.Add(wantDoc)
	check(1, terms.NumUnique())

	terms.Iter(func(term term.T) {
		if !idx.Get(term).Has(wantDoc) {
			t.Fatalf("index should have a doc %v for term %v", wantDoc, term)
		}
	})

}
