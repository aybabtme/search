package query

import (
	"github.com/aybabtme/search/document"
)

// Q is a query, represented by a document
// containing the words to match.
type Q document.Doc
