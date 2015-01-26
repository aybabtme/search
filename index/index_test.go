package index_test

import (
	"github.com/aybabtme/search/document"
	"github.com/aybabtme/search/index"
	"github.com/aybabtme/search/term"
	"math"
	"reflect"
	"testing"
)

func TestHashIndex(t *testing.T) {
	testIndexImpl(t, new(index.HashIndex))
}

func testIndexImpl(t testing.TB, idx index.Idx) {
	check := func(numDocs, numTerms int) {
		if want, got := numDocs, idx.TotalNumDocs(); want != got {
			t.Fatalf("want numdocs %v, got %v", want, got)
		}
		if want, got := numTerms, idx.TotalNumTerms(); want != got {
			t.Fatalf("want numterms %v, got %v", want, got)
		}
	}

	// check that the nil value works
	check(0, 0)
	nosuchterm := term.T("nosuchterm")
	if want, got := 0, idx.Get(nosuchterm).NumDocs(); want != got {
		t.Fatalf("want numdocs %v, got %v", want, got)
	}
	check(0, 0)

	A, B, C := term.T("A"), term.T("B"), term.T("C")

	terms := new(term.HashBag).Add(A, B, C)
	wantDoc := document.NewD(0, terms)

	idx.Add(wantDoc)
	check(1, terms.NumUnique())

	terms.Iter(func(term term.T) {
		if !idx.Get(term).Has(wantDoc) {
			t.Fatalf("index should have a doc %v for term %v", wantDoc, term)
		}
	})

	D := term.T("D")
	moreTerms := new(term.HashBag).Add(B, D)
	anotherDoc := document.NewD(1, moreTerms)

	idx.Add(anotherDoc)
	check(2, 4) // should have 2 documents, 4 terms

	tests := []struct {
		term          term.T
		shouldHave    []document.Doc
		shouldNotHave []document.Doc
	}{
		{A, []document.Doc{wantDoc}, []document.Doc{anotherDoc}},
		{B, []document.Doc{wantDoc, anotherDoc}, []document.Doc{}},
		{C, []document.Doc{wantDoc}, []document.Doc{anotherDoc}},
		{D, []document.Doc{anotherDoc}, []document.Doc{wantDoc}},
	}

	for _, tt := range tests {
		docset := idx.Get(tt.term)
		for _, shouldNotHave := range tt.shouldNotHave {
			if docset.Has(shouldNotHave) {
				t.Fatalf("doc %#v shouldn't contain term %v", docset, tt.term)
			}
		}
		for _, shouldHave := range tt.shouldHave {
			if !docset.Has(shouldHave) {
				t.Fatalf("doc %#v shouldn't contain term %v", docset, tt.term)
			}
		}

		if want, got := len(tt.shouldHave), docset.NumDocs(); want != got {
			t.Fatalf("want numdoc %d, got %d", want, got)
		}
	}

	wantAll := map[int]document.Doc{
		wantDoc.ID():    wantDoc,
		anotherDoc.ID(): anotherDoc,
	}
	gotAll := map[int]document.Doc{}
	idx.Iter(func(doc document.Doc) {
		gotAll[doc.ID()] = doc
	})

	if !reflect.DeepEqual(wantAll, gotAll) {
		t.Logf("want=%#v", wantAll)
		t.Logf(" got=%#v", gotAll)
		t.Fatalf("mismatch!")
	}
}

func repeatTerm(t term.T, frequency int) []term.T {
	terms := make([]term.T, 0, frequency)
	for i := 0; i < frequency; i++ {
		terms = append(terms, t)
	}
	return terms
}

func TestTermFrequency(t *testing.T) {
	tests := []struct {
		term term.T
		freq int
		tf   float64
	}{
		{term: term.T("A"), freq: 3, tf: 3.0 / 3.0},
		{term: term.T("B"), freq: 2, tf: 2.0 / 3.0},
		{term: term.T("C"), freq: 1, tf: 1.0 / 3.0},
	}

	terms := new(term.HashBag)
	doc := document.NewD(0, terms)
	idx := new(index.HashIndex)
	idx.Add(doc)

	// add the terms to the doc
	for _, tt := range tests {
		// term should not be in the doc in advance
		if want, got := 0.0, index.TermFreq(idx, tt.term, doc); want != got {
			t.Fatalf("want tf %f, got %f", want, got)
		}
		terms.Add(repeatTerm(tt.term, tt.freq)...)
	}

	for _, tt := range tests {
		got := index.TermFreq(idx, tt.term, doc)
		if tt.tf != got {
			t.Fatalf("term %v, want tf %f, got %f", tt.term, tt.tf, got)
		}
	}
}

func TestDocumentFrequency(t *testing.T) {
	A := term.T("A")
	B := term.T("B")

	idx := new(index.HashIndex)

	// before adding anything, the df should be 0
	wantdf := 0.0
	gotdf := index.DocumentFreq(idx, A)
	if wantdf != gotdf {
		t.Fatalf("want df %f, got %f", wantdf, gotdf)
	}

	withA := 100
	withoutA := withA * 3

	for i := 0; i < withA+withoutA; i++ {
		terms := new(term.HashBag).Add(B)
		if i < withA {
			terms.Add(A)
		}
		idx.Add(document.NewD(i, terms))
	}

	wantdf = float64(withA)
	gotdf = index.DocumentFreq(idx, A)
	if wantdf != gotdf {
		t.Fatalf("want df %f, got %f", wantdf, gotdf)
	}
}

func TestInvDocumentFrequency(t *testing.T) {
	A := term.T("A")
	B := term.T("B")

	idx := new(index.HashIndex)

	// before adding anything, the df should be 0
	wantidf := math.Inf(+1)
	gotidf := index.InvDocumentFreq(idx, A)
	if wantidf != gotidf {
		t.Fatalf("want idf %f, got %f", wantidf, gotidf)
	}

	withA := 100
	withoutA := withA * 3

	for i := 0; i < withA+withoutA; i++ {
		terms := new(term.HashBag).Add(B)
		if i < withA {
			terms.Add(A)
		}
		idx.Add(document.NewD(i, terms))
	}

	wantdf := float64(withA)
	wantidf = math.Log(float64(idx.TotalNumDocs()) / wantdf)
	gotidf = index.InvDocumentFreq(idx, A)
	if wantidf != gotidf {
		t.Fatalf("want idf %f, got %f", wantidf, gotidf)
	}
}

func TestTFIDF(t *testing.T) {

	A, B, C := term.T("A"), term.T("B"), term.T("C")
	idx := new(index.HashIndex)

	/*
		Given a document containing terms with given frequencies:
			A(3), B(2), C(1)

		Assume collection contains 10,000 documents and document
		frequencies of these terms are:
			A(50), B(1300), C(250)
	*/
	wantA := 50
	wantB := 1300
	wantC := 250
	As, Bs, Cs := repeatTerm(A, wantA), repeatTerm(B, wantB), repeatTerm(C, wantC)

	// make 1 doc with the right frequencies
	doc := document.NewD(0, new(term.HashBag).
		Add(As[:3]...).
		Add(Bs[:2]...).
		Add(Cs[:1]...),
	)
	// use the rest to seed all other docs
	As, Bs, Cs = As[1:], Bs[1:], Cs[1:]

	alldocs := []document.Doc{doc}
	for len(alldocs) < 10000 {
		id := len(alldocs)
		terms := new(term.HashBag)

		// take some terms as long
		// as there are left
		if len(As) > 0 {
			terms.Add(As[0])
			As = As[1:]
		}
		if len(Bs) > 0 {
			terms.Add(Bs[0])
			Bs = Bs[1:]
		}
		if len(Cs) > 0 {
			terms.Add(Cs[0])
			Cs = Cs[1:]
		}
		alldocs = append(alldocs, document.NewD(id, terms))
	}

	idx.Add(alldocs...)

	// verify the test data is as
	// described
	if want, got := wantA, idx.Get(A).NumDocs(); want != got {
		t.Errorf("want A numcount %d, got %d", want, got)
	}
	if want, got := wantB, idx.Get(B).NumDocs(); want != got {
		t.Errorf("want B numcount %d, got %d", want, got)
	}
	if want, got := wantC, idx.Get(C).NumDocs(); want != got {
		t.Errorf("want C numcount %d, got %d", want, got)
	}

	if want, got := 50.0, index.DocumentFreq(idx, A); want != got {
		t.Fatalf("term %q should be in %f docs, was in %f", A, want, got)
	}

	/*
	    Then:
	   		A: tf = 3/3; idf = log(10000/50) = 5.3; tf-idf = 5.3
	   		B: tf = 2/3; idf = log(10000/1300) = 2.0; tf-idf = 1.3
	   		C: tf = 1/3; idf = log(10000/250) = 3.7; tf-idf = 1.2
	*/

	tests := []struct {
		term  term.T
		tf    float64
		idf   float64
		tfidf float64
	}{
		{A, 3.0 / 3.0, math.Log(10000.0 / 50.0), 5.3},
		{B, 2.0 / 3.0, math.Log(10000.0 / 1300.0), 1.36},
		{C, 1.0 / 3.0, math.Log(10000.0 / 250.0), 1.23},
	}

	feq := func(a, b float64) bool {
		diff := math.Abs(a - b)
		magnitudeDiff := diff / math.Max(a, b)
		return magnitudeDiff < 0.01 // less than 1% different
	}

	for _, tt := range tests {
		t.Logf("looking at %v", tt.term)
		if want, got := tt.tf, index.TermFreq(idx, tt.term, doc); !feq(want, got) {
			t.Errorf(" - want TF be %f, got %f", want, got)
		}

		if want, got := tt.idf, index.InvDocumentFreq(idx, tt.term); !feq(want, got) {
			t.Errorf(" - want IDF be %f, got %f", want, got)
		}

		if want, got := tt.tfidf, index.TFIDF(idx, tt.term, doc); !feq(want, got) {
			t.Errorf(" - want TFIDF be %f, got %f", want, got)
		}
	}
}
