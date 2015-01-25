package stem

import "github.com/aybabtme/search/word"

// Stemmer can take words and return their root.
type Stemmer interface {
	Stem(word word.W) (root word.W)
}
