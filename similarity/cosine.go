package similarity

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/query"
	"github.com/aybabtme/search/term"
	"math"
)

// Cosine is apparently better than inner product, they say.
// I haven't read why just yet.
func Cosine(w Weighter, q query.Q, doc document.Doc) float64 {
	innerProduct := InnerProduct(w, q, doc)
	if innerProduct == 0 {
		return 0
	}

	var doclen, qlen float64
	doc.Terms().Iter(func(t term.T) {
		wdoc := w(t, doc)
		wq := w(t, q)
		doclen += wdoc * wdoc
		qlen += wq * wq
	})
	vlen := math.Sqrt(doclen * qlen)

	return innerProduct / vlen
}
