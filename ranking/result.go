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
	Top   int
	Ranks []Scoring
}

func (r *Result) init() {
	if r.Top == 0 {
		r.Top = DefaultTop
	}
	if r.Ranks == nil {
		r.Ranks = make([]Scoring, 0, r.Top)
	}
}

// Add a score/document pair if the score is high enough.
func (r *Result) Add(score float64, doc document.Doc) {
	r.init()
	minScore := r.Ranks[len(r.Ranks)-1].Score
	if score < minScore {
		return // ignore it
	}
	r.Ranks = append(r.Ranks, Scoring{Score: score, Doc: doc})
	sort.Sort(byScore(r.Ranks))
	if len(r.Ranks) > r.Top {
		r.Ranks = r.Ranks[:r.Top]
	}
}

type byScore []Scoring

func (b byScore) Len() int           { return len(b) }
func (b byScore) Less(i, j int) bool { return b[i].Score > b[j].Score }
func (b byScore) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
