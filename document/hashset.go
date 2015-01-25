package document

var _ Set = new(HashSet)

// HashSet is a set implemented with a hash table. The
// nil value is safe to use.
type HashSet struct {
	set map[int]Doc
}

func (h *HashSet) init() {
	h.set = make(map[int]Doc)
}

// Add the doc to the bag.
func (h *HashSet) Add(docs ...Doc) Set {
	if h.set == nil {
		h.init()
	}
	for _, doc := range docs {
		h.set[doc.ID()] = doc
	}
	return h
}

// Has a document or not.
func (h *HashSet) Has(doc Doc) bool {
	if h.set == nil {
		return false
	}
	_, ok := h.set[doc.ID()]
	return ok
}

// Iter iterate over all the documents in the set.
func (h *HashSet) Iter(ƒ func(Doc)) {
	for _, doc := range h.set {
		ƒ(doc)
	}
}

// NumDocs tells how many documents in the set.
func (h *HashSet) NumDocs() int {
	return len(h.set)
}
