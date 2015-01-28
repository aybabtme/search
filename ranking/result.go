package ranking

import (
	"github.com/aybabtme/search/document"
	"sort"
)

// DefaultTop results returned by a ranking.
const DefaultTop = 10

// Scoring of a document.
type Scoring struct {
	Score float64
	Doc   document.Doc
}

// Result of a ranking.
type Result struct {
	sorted bool
	ranks  []Scoring
}

func (r *Result) init() {
	if r.ranks == nil {
		r.ranks = make([]Scoring, 0, DefaultTop)
	}
}

// Rank returns the ordered scored documents, from
// highest score to lowest score.
func (r *Result) Rank() []Scoring {
	if !r.sorted {
		sort.Sort(byScore(r.ranks))
		r.sorted = true
	}
	return r.ranks
}

// Add a score/document pair if the score is high enough.
func (r *Result) Add(score float64, doc document.Doc) {
	r.init()
	r.sorted = false
	r.ranks = append(r.ranks, Scoring{Score: score, Doc: doc})
	sort.Sort(byScore(r.ranks))
}

type byScore []Scoring

func (b byScore) Len() int           { return len(b) }
func (b byScore) Less(i, j int) bool { return b[i].Score > b[j].Score }
func (b byScore) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
