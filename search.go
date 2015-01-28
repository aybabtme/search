package search

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/index"
	"github.com/aybabtme/search/preprocess"
	"github.com/aybabtme/search/query"
	"github.com/aybabtme/search/ranking"
	"github.com/aybabtme/search/similarity"
	"github.com/aybabtme/search/term"
	"io"
)

func tfidfWeigther(idx index.Idx) similarity.Weighter {
	return func(t term.T, d document.Doc) float64 {
		return index.TFIDF(idx, t, d)
	}
}

func dfWeigther(idx index.Idx) similarity.Weighter {
	return func(t term.T, d document.Doc) float64 {
		return float64(d.Terms().Count(t))
	}
}

// DefaultWeighterFactory is used to define the
// similarity.Weighter implementation used by searches.
var DefaultWeighterFactory = dfWeigther

type Search struct {
	Preprocessor   preprocess.Preprocessor
	Idx            index.Idx
	Weigther       similarity.Weighter
	TermBagFactory func() term.Bag
	id             int
}

func (s *Search) init() {
	if s.Preprocessor == nil {
		s.Preprocessor = new(preprocess.English)
	}
	if s.Idx == nil {
		s.Idx = new(index.HashIndex)
	}
	if s.Weigther == nil {
		s.Weigther = DefaultWeighterFactory(s.Idx)
	}
	if s.TermBagFactory == nil {
		s.TermBagFactory = term.DefaultBagFactory
	}
}

// AddReader the content of r to the search index.
func (s *Search) AddReader(r io.Reader) (*Search, error) {
	s.init()
	terms, err := s.Preprocessor.Process(r)
	if err == nil {
		s.Idx.Add(document.NewD(s.id, terms))
		s.id++
	}
	return s, err
}

// AddTerms puts the terms in a document of the search
// index.
func (s *Search) AddTerms(terms term.Bag) *Search {
	s.Idx.Add(document.NewD(s.id, terms))
	s.id++
	return s
}

// QueryReader the search index for an ensemble of terms.
func (s *Search) QueryReader(top int, r io.Reader) ([]ranking.Scoring, error) {
	terms, err := s.Preprocessor.Process(r)
	if err != nil {
		return nil, err
	}
	q := query.Q(document.NewD(-1, terms))
	result := new(ranking.Result)
	s.Idx.Iter(func(doc document.Doc) {
		score := similarity.Cosine(s.Weigther, q, doc)
		result.Add(score, doc)
	})
	return result.Rank()[:top], nil
}

// QueryTerms the search index for an ensemble of terms.
func (s *Search) QueryTerms(top int, terms ...term.T) []ranking.Scoring {
	bag := s.TermBagFactory().Add(terms...)
	q := query.Q(document.NewD(-1, bag))

	result := new(ranking.Result)
	s.Idx.Iter(func(doc document.Doc) {
		score := similarity.Cosine(s.Weigther, q, doc)
		result.Add(score, doc)
	})
	return result.Rank()[:top]
}
