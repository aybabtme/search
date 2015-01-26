package stem_test

import (
	"github.com/aybabtme/search/preprocess/stem"
	"github.com/aybabtme/search/term"
	"testing"
)

func TestEnglishStemmer(t *testing.T) {
	s := new(stem.English)
	tests := []struct {
		from string
		to   string
	}{
		{"accumulation", "accumul"},
		{"elucubration", "elucubr"},
		{"anticonstitutionellement", "anticonstitutionel"},
		{"conspiracy", "conspiraci"},
	}
	for _, tt := range tests {
		want := term.T(tt.to)
		got := s.Stem(term.T(tt.from))
		if want != got {
			t.Fatalf("want %v got %v", want, got)
		}
	}
}
