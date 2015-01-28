package main

import (
	"bytes"
	"encoding/hex"
	"encoding/xml"
	"testing"
	"time"
)

func TestTopicXML(t *testing.T) {
	want := []byte(`<top>
  <num>Number: MB001</num>
  <title>BBC World Service staff cuts</title>
  <querytime>Tue Feb 08 12:30:27 +0000 2011</querytime>
  <querytweettime>34952194402811904</querytweettime>
</top>`)

	qt, err := time.Parse(time.RubyDate, "Tue Feb 08 12:30:27 +0000 2011")
	if err != nil {
		t.Fatal(err)
	}

	resp := NewTopic(
		"Number: MB001",
		"BBC World Service staff cuts",
		qt,
		34952194402811904,
	)

	got, err := xml.MarshalIndent(resp, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Logf("want=%s", hex.Dump(want))
		t.Logf(" got=%s", hex.Dump(got))
		t.Fatal("mismatch")
	}
}

func TestTopicNum(t *testing.T) {

	tests := []struct {
		Number string
		Topic  string
	}{
		{"Number : MB001", "MB001"},
		{"Number: MB001", "MB001"},
		{"Number :MB001", "MB001"},
		{"Number:MB001", "MB001"},
	}
	for _, tt := range tests {
		resp := NewTopic(
			tt.Number,
			"BBC World Service staff cuts",
			time.Now(),
			34952194402811904,
		)

		if resp.TopicNum() != tt.Topic {
			t.Fatalf("want topic %q got %q", tt.Topic, resp.TopicNum())
		}
	}

}
