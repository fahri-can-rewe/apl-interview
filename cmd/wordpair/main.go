package main

import (
	"apl-interview/anagram"
	"apl-interview/httpclient"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

const fiveSecondsTimeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	endpoint := constructEndpoint()
	client := httpclient.NewAPIClient(
		httpclient.WithEndpoint(endpoint),
	)

	var checker anagram.Checker = anagram.FreqMapChecker{}

	if err := run(ctx, os.Args[1:], client, checker); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string, fetcher WordPairFetcher, checker anagram.Checker) error {
	if _, err := parseArgs(args); err != nil {
		return err
	}

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

type WordPairFetcher interface {
	FetchWordPair(ctx context.Context) (*httpclient.WordPair, error)
}

func constructEndpoint() string {
	base, _ := url.Parse(defaultAPIBaseURL)
	endpoint, _ := url.JoinPath(base.String(), apiPath)
	return endpoint
}
