package index

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/term"
	"math"
)

// Idx is an index of terms to document.
type Idx interface {
	Add(term.T, document.Doc)
	Get(term.T) document.Set
	NumDocs() int
}

// TermFreq is the frequency of a term in a specific
// document of the index.
func TermFreq(idx Idx, t term.T, d document.Doc) float64 {
	f := d.Terms().Count(t)
	l := d.Terms().NumTotal()
	return float64(f) / float64(l)
}

// DocumentFreq tells the frequency of a term in all the
// documents of the index.
func DocumentFreq(idx Idx, t term.T) float64 {
	withTerm := idx.Get(t).NumDocs()
	allDocs := idx.NumDocs()
	return float64(withTerm) / float64(allDocs)
}

// InvDocumentFreq is the inverse document frequency of
// a term in the documents of the index, dampened by
// log2.
func InvDocumentFreq(idx Idx, t term.T) float64 {
	N := float64(idx.NumDocs())
	df := DocumentFreq(idx, t)
	return math.Log2(N / df)
}

// TFIDF is the score of a term given its frequency in
// a document, weighted relatively to its frequency in
// all documents.
//
// This means that terms appearing in all documents (more
// generic terms) of the index will score less than those
// occuring in a few of them (more specialized terms).
func TFIDF(idx Idx, t term.T, d document.Doc) float64 {
	tf := TermFreq(idx, t, d)
	idf := InvDocumentFreq(idx, t)
	return tf * idf
}
