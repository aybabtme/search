package similarity_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/query"
	"github.com/aybabtme/search/similarity"
	"github.com/aybabtme/search/term"
	"math"
	"testing"
)

func feq(a, b float64) bool {
	diff := math.Abs(a - b)
	if diff == 0 {
		return true
	}
	magnitudeDiff := diff / math.Max(a, b)
	return magnitudeDiff < 0.01 // less than 1% different
}

func TestCosine(t *testing.T) {
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

	if want, got := 0.81, similarity.Cosine(w, q, doc1); !feq(want, got) {
		t.Errorf("want similarity doc1 %f, got %f", want, got)
	}

	if want, got := 0.13, similarity.Cosine(w, q, doc2); !feq(want, got) {
		t.Fatalf("want similarity doc2 %f, got %f", want, got)
	}
}

func TestCosineNoMatch(t *testing.T) {
	a, b, c := term.T("1"), term.T("2"), term.T("3")

	doc1 := document.NewD(1, new(term.HashBag).Add(
		a, a,
		b, b, b,
	))
	q := query.Q(document.NewD(3, new(term.HashBag).Add(
		c, c,
	)))

	w := similarity.Weighter(func(term term.T, doc document.Doc) float64 {
		return float64(doc.Terms().Count(term))
	})

	if want, got := 0.0, similarity.Cosine(w, q, doc1); !feq(want, got) {
		t.Errorf("want similarity doc1 %f, got %f", want, got)
	}
}
