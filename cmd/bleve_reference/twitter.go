package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"github.com/aybabtme/uniplot/spark"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"github.com/davecheney/profile"
	"io"
	"log"
	"os"
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

func createFileOrExit(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		printUsage("can't create %q: %v", filename, err)
	}
	return f
}

func main() {
	p := profile.Start(&profile.Config{
		CPUProfile: true,
	})
	defer p.Stop()

	log.SetFlags(0)
	corpusFilename := flag.String("corpus", "", "file containing the tweets")
	trecFilename := flag.String("trec", "", "file containing the trec queries")
	top := flag.Int("top", 1000, "return top N results")
	outputFilename := flag.String("output", "Result.txt", "file where to write the trec results")
	tag := flag.String("tag", "antoineRun", "tag to append to the end of the trec results")
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
	output := createFileOrExit(*outputFilename)
	defer output.Close()

	log.SetPrefix("twitter_trec: ")

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("", mapping)
	if err != nil {
		log.Fatalf("creating index file: %v", err)
	}

	idToTweet, err := indexTweets(index, corpus)
	if err != nil {
		log.Fatalf("indexing %q: %v", corpus.Name(), err)
	}

	result := log.New(output, "", 0)
	report := func(t Topic, hits search.DocumentMatchCollection) {
		for k, r := range hits {
			tweet, ok := idToTweet[r.ID]
			if !ok {
				log.Fatalf("doesn't exist: tweet ID %q", r.ID)
			}
			result.Printf("%s Q0 %s %d %.3f %s",
				t.Number,
				tweet.ID,
				k+1,
				r.Score,
				*tag,
			)
		}
	}

	if err := answerQueries(index, *top, trecQueries, report); err != nil {
		log.Fatalf("answering queries in %q: %v", trecQueries.Name(), err)
	}
}

func indexTweets(idx bleve.Index, corpus io.Reader) (map[string]Tweet, error) {
	log.Printf("decoding and indexing tweets...")
	tweets, errc := decodeTweets(corpus)

	sparkUI := spark.Spark(time.Millisecond * 23)
	sparkUI.Units = " tweets"
	sparkUI.Start()
	defer sparkUI.Stop()

	idToTweet := map[string]Tweet{}
	for t := range tweets {
		sparkUI.Add(1)
		idx.Index(t.ID, t)
		idToTweet[t.ID] = t
	}
	log.Printf("indexed %d tweets", len(idToTweet))
	return idToTweet, <-errc
}

func answerQueries(idx bleve.Index, top int, trecQueries io.Reader, fn func(t Topic, hits search.DocumentMatchCollection)) error {
	log.Printf("decoding and answering queries...")
	topics, errc := decodeTopicQueries(trecQueries)

	sparkUI := spark.Spark(time.Millisecond * 197)
	sparkUI.Units = " topics"
	sparkUI.Start()
	defer sparkUI.Stop()

	i := 0
	for t := range topics {

		sparkUI.Add(1)

		query := bleve.NewQueryStringQuery(t.Title)
		req := bleve.NewSearchRequest(query)
		req.Size = top
		res, err := idx.Search(req)
		if err != nil {
			return err
		}
		i++
		fn(t, res.Hits)
	}
	log.Printf("answered %d queries", i)
	return <-errc
}

func decodeTweets(r io.Reader) (<-chan Tweet, <-chan error) {
	tweets := make(chan Tweet)
	errc := make(chan error, 1)

	go func() {
		defer close(errc)
		defer close(tweets)

		scan := bufio.NewScanner(r)
		scan.Split(bufio.ScanLines)

		for scan.Scan() {
			t := Tweet{}
			err := t.Decode(scan.Bytes())
			if err != nil {
				errc <- err
				return
			}
			tweets <- t
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
		dec := xml.NewDecoder(r)
		for {
			t := Topic{}
			err := dec.Decode(&t)
			if err == nil {
				topics <- t
			} else if err != io.EOF {
				errc <- err
				return
			} else {
				return
			}
		}
	}()

	return topics, errc
}
