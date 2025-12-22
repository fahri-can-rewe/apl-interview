package main

import (
	"apl-interview/anagram"
	"apl-interview/httpclient"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/url"
)

const (
	defaultAPIBaseURL = "https://interview.sowula.at"
	apiPath           = "/word-pair"
)

type config struct {
	APIBaseURL string
}

func parseArgs(args []string) (config, error) {
	fs := flag.NewFlagSet("wordpair", flag.ContinueOnError)
	apiBaseURL := fs.String("apiBaseUrl", defaultAPIBaseURL, "Base URL of the Word-Pair API")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	return config{APIBaseURL: *apiBaseURL}, nil
}

func buildAPIClient(apiBaseURL string) (*httpclient.APIClient, error) {
	address, err := url.Parse(apiBaseURL)
	if err != nil || address.Scheme == "" || address.Host == "" {
		return nil, fmt.Errorf("bad --apiBaseUrl: %q", apiBaseURL)
	}
	endpoint, _ := url.JoinPath(address.String(), apiPath)
	client := httpclient.NewAPIClient(httpclient.WithEndpoint(endpoint))
	return client, nil
}

func fetchWordPair(ctx context.Context, ac *httpclient.APIClient) (*httpclient.WordPair, error) {
	wp, err := ac.FetchWordPair(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get word pair: %w", err)
	}
	return wp, nil
}

func validateWordPair(word1, word2 string) error {
	if !anagram.IsAlphabetic(word1) || !anagram.IsAlphabetic(word2) {
		return errors.New("words must contain only letters")
	}
	if len([]rune(word1)) != len([]rune(word2)) {
		return fmt.Errorf("words do not match in length, word1: %q & word2: %q", word1, word2)
	}
	return nil
}

func areAnagrams(word1, word2 string, strategy anagram.Checker) bool {
	return strategy.AreAnagrams(word1, word2)
}
