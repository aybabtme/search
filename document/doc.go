package document

import (
	"github.com/aybabtme/search/term"
)

// Doc is a document.
type Doc interface {
	// ID must be unique to this document.
	ID() int
	// Terms in the document.
	Terms() term.Bag
	// Value is the original document.
	Value() interface{}
}

// Set of documents.
type Set interface {
	// Add the docs to the bag.
	Add(docs ...Doc) Set
	// Has a document or not.
	Has(doc Doc) bool
	// NumDocs tells how many documents in the set.
	NumDocs() int
	// Iter iterates over all the documents in the set.
	Iter(func(Doc))
}
