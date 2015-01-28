package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"github.com/aybabtme/search"
	"github.com/aybabtme/uniplot/spark"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func printUsage(format string, args ...interface{}) {
	log.Printf(format, args...)
	flag.PrintDefaults()
	os.Exit(1)
}

func openFileOrExit(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		printUsage("can't open %q: %v", filename, err)
	}
	return f
}

func main() {
	log.SetFlags(0)
	corpusFilename := flag.String("corpus", "", "file containing the tweets")
	trecFilename := flag.String("trec", "", "file containing the trec queries")
	top := flag.Int("top", 1000, "return top N results")
	flag.Parse()

	switch {
	case *corpusFilename == "":
		printUsage("need to specify a corpus")
	case *trecFilename == "":
		printUsage("need to specify a corpus")
	}

	corpus := openFileOrExit(*corpusFilename)
	defer corpus.Close()
	trecQueries := openFileOrExit(*trecFilename)
	defer trecQueries.Close()

	log.SetPrefix("twitter_trec: ")

	s := new(search.Search)
	if err := indexTweets(s, corpus); err != nil {
		log.Fatalf("indexing %q: %v", corpus.Name(), err)
	}

	if err := answerQueries(s, *top, trecQueries); err != nil {
		log.Fatalf("answering queries in %q: %v", trecQueries.Name(), err)
	}
}

func indexTweets(s *search.Search, corpus io.Reader) error {
	log.Printf("decoding and indexing tweets...")
	tweets, errc := decodeTweets(corpus)

	sparkUI := spark.Spark(time.Millisecond * 23)
	sparkUI.Units = " tweets"
	sparkUI.Start()
	defer sparkUI.Stop()

	count := 0
	for t := range tweets {
		sparkUI.Add(1)
		s.AddReader(t)
		count++
	}
	log.Printf("indexed %d tweets", count)
	return <-errc
}

func answerQueries(s *search.Search, top int, trecQueries io.Reader) error {
	log.Printf("decoding and answering queries...")
	topics, errc := decodeTopicQueries(trecQueries)

	sparkUI := spark.Spark(time.Millisecond * 197)
	sparkUI.Units = "topics"
	sparkUI.Start()
	defer sparkUI.Stop()

	for t := range topics {
		sparkUI.Add(1)

		ranks, err := s.QueryReader(top, strings.NewReader(t.Title))
		if err != nil {
			return err
		}
		_ = ranks
	}
	return <-errc
}

func decodeTweets(r io.Reader) (<-chan io.Reader, <-chan error) {
	tweets := make(chan io.Reader)
	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		defer close(tweets)

		scan := bufio.NewScanner(r)
		scan.Split(bufio.ScanLines)

		for scan.Scan() {
			tweets <- strings.NewReader(scan.Text())
		}

		if scan.Err() != nil {
			errc <- scan.Err()
		}
	}()

	return tweets, errc
}

func decodeTopicQueries(r io.Reader) (<-chan Topic, <-chan error) {
	topics := make(chan Topic)
	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		defer close(topics)

		var err error
		dec := xml.NewDecoder(r)
		for err == nil {
			t := Topic{}
			err = dec.Decode(&t)
			if err == nil {
				topics <- t
			}
		}

		if err != io.EOF {
			errc <- err
		}
	}()

	return topics, errc
}
