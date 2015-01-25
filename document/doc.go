package document

import (
	"github.com/aybabtme/search/term"
)

// Doc is a document.
type Doc interface {
	ID() int
	Terms() term.Bag
}

// Set of documents.
type Set interface {
	// Add the docs to the bag.
	Add(docs ...Doc)
	// Has a document or not.
	Has(doc Doc) bool
	// NumDocs tells how many documents in the set.
	NumDocs() int
}
