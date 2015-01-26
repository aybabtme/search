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
}

func (h *HashIndex) init() {
	if h.DocsetFactory == nil {
		h.DocsetFactory = DefaultDocsetFactory
	}
	if h.mapping == nil {
		h.mapping = make(map[term.T]document.Set)
	}
	// ATTENTION: _must_ be done after nil-test of `DocsetFactory`
	if h.allDocs == nil {
		h.allDocs = h.DocsetFactory()
	}
}

// Add indexes the documents' terms.
func (h *HashIndex) Add(docs ...document.Doc) Idx {
	h.init()
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
