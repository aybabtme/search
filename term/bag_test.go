package term_test

import (
	"github.com/aybabtme/search/term"
	"reflect"
	"testing"
)

func TestHashBag(t *testing.T) {
	testBagImpl(t, term.DefaultBagFactory())
}

func testBagImpl(t testing.TB, bag term.Bag) {

	bag.Iter(func(tt term.T) {
		t.Fatalf("bag already contains a term (%v): %#v", tt, bag)
	})

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

	// add it twice
	bag.Add(another).Add(another)
	check(another, 3, 2, wantFreq+3)

	wantAll := map[term.T]struct{}{
		wantWord: struct{}{},
		another:  struct{}{},
	}
	gotAll := make(map[term.T]struct{})
	bag.Iter(func(t term.T) {
		gotAll[t] = struct{}{}
	})
	if !reflect.DeepEqual(wantAll, gotAll) {
		t.Logf("want=%#v", wantAll)
		t.Logf(" got=%#v", gotAll)
		t.Fatal("mismatch between content")
	}
}
