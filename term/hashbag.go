package term

var _ Bag = new(HashBag)

// HashBag is a bag implemented with a hash table. The
// nil value is safe to use.
type HashBag struct {
	size int
	bag  map[T]int
}

func (h *HashBag) init() {
	h.bag = make(map[T]int)
}

// Add the term to the bag.
func (h *HashBag) Add(terms ...T) Bag {
	if h.bag == nil {
		h.init()
	}
	for _, t := range terms {
		h.size++
		h.bag[t] = h.bag[t] + 1
	}
	return h
}

// Count is the number of occurence of the term
// in a bag.
func (h *HashBag) Count(term T) int {
	if h.bag == nil {
		return 0
	}
	return h.bag[term]
}

// NumUnique is the number of unique terms in the bag.
func (h *HashBag) NumUnique() int {
	return len(h.bag)
}

// NumTotal is the number of terms in the bag, counting
// duplicates.
func (h *HashBag) NumTotal() int {
	return h.size
}

// Iter iterates over all unique terms in the bag.
func (h *HashBag) Iter(ƒ func(T)) {
	for t := range h.bag {
		ƒ(t)
	}
}
