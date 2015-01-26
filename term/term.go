package term

// TODO(antoine): memory optimization, make
// weak references to term, keep a repository
// of terms with the actual memory.

// T is a term.
type T string

// Bag of terms.
type Bag interface {
	// Add the terms to the bag.
	Add(terms ...T) Bag
	// Count how many times the term is found in the bag.
	Count(term T) int
	// NumUnique terms in the bag.
	NumUnique() int
	// NumTotal is how many terms in the bag, counting duplicates.
	NumTotal() int
	// Iter iterates over all unique terms in the bag.
	Iter(func(T))
}

// DefaultBagFactory create using a HashBag implementation.
func DefaultBagFactory() Bag { return new(HashBag) }
