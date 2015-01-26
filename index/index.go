package index

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/term"
	"math"
)

// Idx is an index of terms to document.
type Idx interface {
	// Add the documents' terms to the index.
	Add(...document.Doc) Idx
	// Get the set of document containing a term.
	Get(term.T) document.Set
	// TotalNumTerms is the total number of terms in the index.
	TotalNumTerms() int
	// TotalNumDocs is the total number of docs in the index.
	TotalNumDocs() int
}

// TermFreq is the frequency of a term in a specific
// document of the index, compared to the total number
// of unique terms in the document.
func TermFreq(idx Idx, t term.T, d document.Doc) float64 {
	l := d.Terms().NumUnique()
	if l == 0 {
		return 0
	}
	f := d.Terms().Count(t)
	return float64(f) / float64(l)
}

// DocumentFreq tells the frequency of a term in all the
// documents of the index.
func DocumentFreq(idx Idx, t term.T) float64 {
	return float64(idx.Get(t).NumDocs())
}

// InvDocumentFreq is the inverse document frequency of
// a term in the documents of the index, dampened by
// log2.
func InvDocumentFreq(idx Idx, t term.T) float64 {
	N := float64(idx.TotalNumDocs())
	df := DocumentFreq(idx, t)
	if df == 0.0 {
		// if the term is never there, then its idf is
		// +inf. This way if its only in 1 document, its
		// importance will be infinite compared to others
		return math.Inf(+1)
	}
	return math.Log(N / df)
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
