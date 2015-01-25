package document

var _ Set = new(HashSet)

// HashSet is a set implemented with a hash table. The
// nil value is safe to use.
type HashSet struct {
	set map[int]struct{}
}

func (h *HashSet) init() {
	h.set = make(map[int]struct{})
}

// Add the doc to the bag.
func (h *HashSet) Add(docs ...Doc) {
	if h.set == nil {
		h.init()
	}
	for _, doc := range docs {
		h.set[doc.ID()] = struct{}{}
	}
}

// Has a document or not.
func (h *HashSet) Has(doc Doc) bool {
	if h.set == nil {
		return false
	}
	_, ok := h.set[doc.ID()]
	return ok
}

// NumDocs tells how many documents in the set.
func (h *HashSet) NumDocs() int {
	return len(h.set)
}
