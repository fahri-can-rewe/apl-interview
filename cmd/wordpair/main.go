package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fahri-can-rewe/apl-interview/internal/anagram"
	"github.com/fahri-can-rewe/apl-interview/internal/config"
	"github.com/fahri-can-rewe/apl-interview/internal/httpclient"
)

const fiveSecondsTimeout = 5 * time.Second

type WordPairFetcher interface {
	FetchWordPair(ctx context.Context) (*httpclient.WordPair, error)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	conf, err := config.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	endpoint, err := config.MakeEndpoint(conf.APIBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	client := httpclient.NewAPIClient(httpclient.WithEndpoint(endpoint))

	var checker anagram.Checker = anagram.SortChecker{}

	if err := run(ctx, client, checker); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, fetcher WordPairFetcher, checker anagram.Checker) error {
	wp, err := fetcher.FetchWordPair(ctx)
	if err != nil {
		return fmt.Errorf("failed to get word pair: %w", err)
	}

	isAnagram := checker.AreAnagrams(wp.FirstWord, wp.SecondWord)

	fmt.Printf("Word 1: %s\n", wp.FirstWord)
	fmt.Printf("Word 2: %s\n", wp.SecondWord)
	fmt.Printf("Are Anagrams: %v\n", isAnagram)
	return nil
}
