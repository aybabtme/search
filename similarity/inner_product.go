package similarity

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/query"
	"github.com/aybabtme/search/term"
	"log"
)

// Weighter gives a weight to a term in a document.
type Weighter func(term.T, document.Doc) float64

// InnerProduct is a simple similarity measurement method. It
// favors long documents with a large number of unique terms
// and measure how many terms are matched, but not how many
// terms are _not_ matched.
//
// The measurement is unbounded.
func InnerProduct(w Weighter, q query.Q, doc document.Doc) float64 {
	var product float64
	i := 0
	q.Terms().Iter(func(t term.T) {
		wij := w(t, doc)
		wiq := w(t, q)
		log.Printf("%d: t=%v, wij=%f, wiq=%f", i, t, wij, wiq)
		product += wij * wiq
		i++
	})
	return product
}
