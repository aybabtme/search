package similarity_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/query"
	"github.com/aybabtme/search/similarity"
	"github.com/aybabtme/search/term"
	"testing"
)

func TestInnerProduct(t *testing.T) {
	a, b, c := term.T("1"), term.T("2"), term.T("3")

	doc1 := document.NewD(1, new(term.HashBag).Add(
		a, a,
		b, b, b,
		c, c, c, c, c,
	))
	doc2 := document.NewD(2, new(term.HashBag).Add(
		a, a, a,
		b, b, b, b, b, b, b,
		c,
	))
	q := query.Q(document.NewD(3, new(term.HashBag).Add(
		c, c,
	)))

	w := similarity.Weighter(func(term term.T, doc document.Doc) float64 {
		return float64(doc.Terms().Count(term))
	})

	if want, got := 10.0, similarity.InnerProduct(w, q, doc1); want != got {
		t.Fatalf("want similarity doc1 %f, got %f", want, got)
	}

	if want, got := 2.0, similarity.InnerProduct(w, q, doc2); want != got {
		t.Fatalf("want similarity doc2 %f, got %f", want, got)
	}
}
