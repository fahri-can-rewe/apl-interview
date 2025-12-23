package main

import (
	"apl-interview/anagram"
	"apl-interview/httpclient"
	"apl-interview/internal"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const fiveSecondsTimeout = 5 * time.Second

type WordPairFetcher interface {
	FetchWordPair(ctx context.Context) (*httpclient.WordPair, error)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	conf, err := internal.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	client, err := internal.BuildAPIClient(conf.APIBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	var checker anagram.Checker = anagram.FreqMapChecker{}

	if err := run(ctx, client, checker); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, fetcher WordPairFetcher, checker anagram.Checker) error {
	wp, err := fetcher.FetchWordPair(ctx)
	if err != nil {
		return fmt.Errorf("failed to get word pair: %w", err)
	}

	if err := anagram.ValidateWordPair(wp.FirstWord, wp.SecondWord); err != nil {
		return err
	}

	isAnagram := checker.AreAnagrams(wp.FirstWord, wp.SecondWord)

	fmt.Printf("Word 1: %s\n", wp.FirstWord)
	fmt.Printf("Word 2: %s\n", wp.SecondWord)
	fmt.Printf("Are Anagrams: %v\n", isAnagram)
	return nil
}
