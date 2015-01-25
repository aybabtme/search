package document

import (
	"github.com/aybabtme/search/term"
)

// D is a document.
type D struct {
	id    int
	terms term.Bag
}

// NewD creates a document.
func NewD(id int, terms term.Bag) *D {
	return &D{id: id, terms: terms}
}

func (d *D) ID() int         { return d.id }
func (d *D) Terms() term.Bag { return d.terms }
