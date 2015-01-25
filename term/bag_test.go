package term_test

import (
	"github.com/aybabtme/search/term"
	"testing"
)

func TestHashBag(t *testing.T) {
	testBagImpl(t, new(term.HashBag))
}

func testBagImpl(t testing.TB, bag term.Bag) {
	check := func(term term.T, freq, unique, total int) {
		if want, got := freq, bag.Count(term); want != got {
			t.Fatalf("want freq %d got %d", want, got)
		}

		if want, got := total, bag.NumTotal(); want != got {
			t.Fatalf("want total %d got %d", want, got)
		}

		if want, got := unique, bag.NumUnique(); want != got {
			t.Fatalf("want unique %d got %d", want, got)
		}
	}

	wantFreq := 3
	wantWord := term.T("hello")

	check(wantWord, 0, 0, 0)

	for i := 0; i < wantFreq; i++ {
		bag.Add(wantWord)
		check(wantWord, i+1, 1, i+1)
	}

	check(wantWord, wantFreq, 1, wantFreq)

	another := term.T("another")
	bag.Add(another)

	check(another, 1, 2, wantFreq+1)
}
