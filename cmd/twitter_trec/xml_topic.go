package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Topic struct {
	Number         string
	Title          string
	QueryTime      time.Time
	QueryTweetTime int
}

func NewTopic(num, title string, queryTime time.Time, queryTweetTime int) *Topic {
	return &Topic{
		Number:         strings.TrimSpace(num),
		Title:          strings.TrimSpace(title),
		QueryTime:      queryTime,
		QueryTweetTime: queryTweetTime,
	}
}

func (t *Topic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type uglyXML struct {
		XMLName        xml.Name `xml:"top"`
		Number         string   `xml:"num"`
		Title          string   `xml:"title"`
		QueryTime      string   `xml:"querytime"`
		QueryTweetTime string   `xml:"querytweettime"`
	}
	var err error
	v := &uglyXML{}
	if err := d.DecodeElement(v, &start); err != nil {
		return err
	}

	t.Number = strings.NewReplacer(
		"Number:", "",
		"Number :", "",
		" ", "",
	).Replace(v.Number)
	t.Title = v.Title
	t.QueryTime, err = time.Parse(time.RubyDate, strings.TrimSpace(v.QueryTime))
	if err != nil {
		return fmt.Errorf("query time was %q: %v", v.QueryTime, err)
	}
	t.QueryTweetTime, err = strconv.Atoi(strings.TrimSpace(v.QueryTweetTime))
	return err
}
