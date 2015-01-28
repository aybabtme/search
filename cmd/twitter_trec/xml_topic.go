package main

import (
	"encoding/xml"
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
		QueryTime:      queryTime.Format(time.RubyDate),
		QueryTweetTime: queryTweetTime,
	}
}

func (t *Topic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type uglyXml struct {
		XMLName        xml.Name `xml:"top"`
		Number         string   `xml:"num"`
		Title          string   `xml:"title"`
		QueryTime      string   `xml:"querytime"`
		QueryTweetTime string   `xml:"querytweettime"`
	}
	var err error
	v := &uglyXml{}
	if err := d.Decode(v); err != nil {
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
		return err
	}
	t.QueryTweetTime, err = strconv.Atoi(strings.TrimSpace(v.Number))
	return err
}

func (t Topic) TopicNum() string {
	return
}
