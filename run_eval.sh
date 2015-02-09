#!/bin/bash

set -e

GOPATH="$HOME/gocode/"


TWITTER_PATH="cmd/twitter_trec"
BLEVE_PATH="cmd/bleve_reference"
TESTDATA_PATH="third_party/testdata"

OWN_COMMAND="go run $TWITTER_PATH/tweet.go $TWITTER_PATH/twitter.go $TWITTER_PATH/xml_topic.go"
BLEVE_COMMAND="go run $BLEVE_PATH/tweet.go $BLEVE_PATH/twitter.go $BLEVE_PATH/xml_topic.go"

COMMAND=$BLEVE_COMMAND
# COMMAND=$OWN_COMMAND

# Input
CORPUS_FILENAME="$TESTDATA_PATH/corpus.txt"
QUERY_FILENAME="$TESTDATA_PATH/trec-queries.txt"
TREC_TRUTH_FILENAME="Trec_microblog11-qrels.txt"

# Midway
RESULT_FILENAME="result.txt"

# Output
REPO_VERSION=`git rev-parse --verify --short HEAD`
SCORE_FILENAME="$REPO_VERSION.score"

FLAGS="-corpus $CORPUS_FILENAME -trec $QUERY_FILENAME -output $RESULT_FILENAME"

RUNNABLE="$COMMAND $FLAGS"

eval $RUNNABLE

./trec_eval $TREC_TRUTH_FILENAME $RESULT_FILENAME > $SCORE_FILENAME
cat $SCORE_FILENAME
