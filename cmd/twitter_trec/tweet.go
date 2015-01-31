package main

import (
	"fmt"
	"regexp"
)

var tweetRe = regexp.MustCompile(`^\s*(\d+)\s*(.*)$`)

type Tweet struct {
	ID      int
	Message string
}

func (t *Tweet) Decode(b []byte) error {

	// skip stuff that's not looking like a number
	var j int
	for i, p := range b {
		if p >= '0' && p <= '9' {
			j = i
			break
		}
	}

	matches := tweetRe.FindSubmatch(b[j:])
	if len(matches) != 3 {
		return fmt.Errorf("need 2 matches, got %d", len(matches))
	}

	for _, num := range matches[1] {
		c := int(num - '0')
		t.ID = (10 * t.ID) + c
	}
	t.Message = string(matches[2])
	return nil
}

func (t Tweet) String() string {
	return fmt.Sprintf("[%d] %s", t.ID, t.Message)
}
