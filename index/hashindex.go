package index

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/term"
)

// DefaultDocsetFactory creates document.Set
// implemented with a document.HashSet.
func DefaultDocsetFactory() document.Set {
	return new(document.HashSet)
}

// HashIndex is an index implemented with a hash table.
// It tracks which terms are present in which documents.
// Documents are stored in sets, which default implementation
// is based on a document.HashSet.
//
// nil value is safe for use.
type HashIndex struct {
	// DocsetFactory is the func used to create new
	// document.Set. If this is nil, it defaults to
	// using DefaultDocsetFactory.
	DocsetFactory func() document.Set

	allDocs document.Set
	mapping map[term.T]document.Set

	cacheDocFreq    map[term.T]float64
	cacheInvDocFreq map[term.T]float64
	cacheTermFreq   map[termDocKey]float64
	cacheTfidf      map[termDocKey]float64
}

type termDocKey struct {
	t term.T
	d int
}

func (h *HashIndex) clearCaches() {
	h.cacheDocFreq = make(map[term.T]float64)
	h.cacheTermFreq = make(map[termDocKey]float64)
	h.cacheInvDocFreq = make(map[term.T]float64)
	h.cacheTfidf = make(map[termDocKey]float64)
}

func (h *HashIndex) init() {
	if h.DocsetFactory == nil {
		h.DocsetFactory = DefaultDocsetFactory
	}
	if h.mapping == nil {
		h.mapping = make(map[term.T]document.Set)
	}

	if h.cacheDocFreq == nil {
		h.cacheDocFreq = make(map[term.T]float64)
	}
	if h.cacheTermFreq == nil {
		h.cacheTermFreq = make(map[termDocKey]float64)
	}
	if h.cacheInvDocFreq == nil {
		h.cacheInvDocFreq = make(map[term.T]float64)
	}
	if h.cacheTfidf == nil {
		h.cacheTfidf = make(map[termDocKey]float64)
	}

	// ATTENTION: _must_ be done after nil-test of `DocsetFactory`
	if h.allDocs == nil {
		h.allDocs = h.DocsetFactory()
	}
}

// Add indexes the documents' terms.
func (h *HashIndex) Add(docs ...document.Doc) Idx {
	h.init()
	h.clearCaches()
	for _, doc := range docs {
		doc.Terms().Iter(func(t term.T) {
			docset, ok := h.mapping[t]
			if !ok {
				docset = h.DocsetFactory()
				h.mapping[t] = docset
			}
			docset.Add(doc)
		})

	}
	h.allDocs.Add(docs...)

	return h
}

// Get the set of document containing a term.
func (h *HashIndex) Get(t term.T) document.Set {
	h.init()
	docset := h.mapping[t]
	if docset == nil {
		// term could be in no documents, return
		// an empty set of documents.
		docset = h.DocsetFactory()
	}
	return docset
}

// TotalNumTerms is the total number of docs in the index.
func (h *HashIndex) TotalNumTerms() int {
	return len(h.mapping)
}

// TotalNumDocs is the total number of docs in the index.
func (h *HashIndex) TotalNumDocs() int {
	h.init()
	return h.allDocs.NumDocs()
}

// Iter iterates over all the documents in the index.
func (h *HashIndex) Iter(ƒ func(document.Doc)) {
	h.init()
	h.allDocs.Iter(ƒ)
}

func (h *HashIndex) TermFreq(t term.T, d document.Doc) float64 {
	key := termDocKey{t, d.ID()}
	if v, ok := h.cacheTermFreq[key]; ok {
		return v
	}
	v := TermFreq(h, t, d)
	h.cacheTfidf[key] = v
	return v
}

func (h *HashIndex) DocumentFreq(t term.T) float64 {
	if v, ok := h.cacheDocFreq[t]; ok {
		return v
	}
	v := DocumentFreq(h, t)
	h.cacheDocFreq[t] = v
	return v
}

func (h *HashIndex) InvDocumentFreq(t term.T) float64 {
	if v, ok := h.cacheInvDocFreq[t]; ok {
		return v
	}
	v := InvDocumentFreq(h, t)
	h.cacheInvDocFreq[t] = v
	return v
}

func (h *HashIndex) TFIDF(t term.T, d document.Doc) float64 {
	key := termDocKey{t, d.ID()}
	if v, ok := h.cacheTfidf[key]; ok {
		return v
	}
	v := TFIDF(h, t, d)
	h.cacheTfidf[key] = v
	return v
}
