package document

import (
	"github.com/aybabtme/search/term"
)

// D is a document.
type D struct {
	id    int
	terms term.Bag
	value interface{}
}

// NewD creates a document.
func NewD(id int, terms term.Bag, v interface{}) *D {
	return &D{id: id, terms: terms, value: v}
}

func (d *D) ID() int            { return d.id }
func (d *D) Terms() term.Bag    { return d.terms }
func (d *D) Value() interface{} { return d.value }
