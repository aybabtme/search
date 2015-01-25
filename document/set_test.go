package document_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/term"
	"reflect"
	"testing"
)

func TestHashSet(t *testing.T) {
	testSetImpl(t, new(document.HashSet))
}

func testSetImpl(t testing.TB, set document.Set) {

	set.Iter(func(doc document.Doc) {
		t.Fatalf("set already contains a doc (%v): %#v", doc, set)
	})

	check := func(document document.Doc, has bool, total int) {
		if want, got := has, set.Has(document); want != got {
			t.Fatalf("want has %v, got %v", want, got)
		}

		if want, got := total, set.NumDocs(); want != got {
			t.Fatalf("want num docs %d got %d", want, got)
		}
	}

	wantDoc := document.NewD(0, new(term.HashBag))
	check(wantDoc, false, 0)

	set.Add(wantDoc)
	check(wantDoc, true, 1)

	// adding doc twice is indepotent
	set.Add(wantDoc)
	check(wantDoc, true, 1)

	// adding a term to a doc doesn't change the set
	wantDoc.Terms().Add(term.T("hello"))
	check(wantDoc, true, 1)

	others := []document.Doc{
		document.NewD(1, new(term.HashBag)),
		document.NewD(2, new(term.HashBag)),
		document.NewD(3, new(term.HashBag)),
	}

	for _, doc := range others {
		check(doc, false, 1)
	}

	set.Add(others...)
	check(wantDoc, true, 1+len(others))
	for _, doc := range others {
		check(doc, true, 1+len(others))
	}

	// add with the chain
	set.Add(others...).Add(others...)
	check(wantDoc, true, 1+len(others))
	for _, doc := range others {
		check(doc, true, 1+len(others))
	}

	var wantAll []document.Doc
	wantAll = append(wantAll, wantDoc)
	wantAll = append(wantAll, others...)
	gotAll := make([]document.Doc, len(wantAll))

	set.Iter(func(doc document.Doc) { gotAll[doc.ID()] = doc })

	if !reflect.DeepEqual(wantAll, gotAll) {
		t.Logf("want=%#v", wantAll)
		t.Logf(" got=%#v", gotAll)
		t.Fatal("mismatch between content")
	}
}
