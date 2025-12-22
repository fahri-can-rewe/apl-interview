package main

import (
	"apl-interview/anagram"
	"apl-interview/httpclient"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	defaultAPIBaseURL = "https://interview.sowula.at"
	apiPath           = "/word-pair"
	reqTimeout        = 5 * time.Second
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
	endpoint := address.ResolveReference(&url.URL{Path: path.Join(address.Path, apiPath)}).String()
	httpClient := &http.Client{Timeout: reqTimeout}
	return httpclient.NewAPIClientWithHTTP(httpClient, endpoint), nil
}

func fetchWordPair(apiClient *httpclient.APIClient) (*httpclient.WordPair, error) {
	wordPair, err := apiClient.FetchWordPair()
	if err != nil {
		return nil, fmt.Errorf("failed to get word pair: %w", err)
	}
	return wordPair, nil
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

func areAnagrams(word1, word2 string) bool {
	return anagram.AreAnagrams(anagram.CountLetters(word1), anagram.CountLetters(word2))
}
