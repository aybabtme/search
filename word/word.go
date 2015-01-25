package word

// W is a word.
type W []rune

// Set of words.
type Set interface {
	Add(word W)
	Has(word W) bool
	All() []W
	Iter(func(W) bool)
}
