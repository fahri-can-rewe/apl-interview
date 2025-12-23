package main

import (
	"apl-interview/httpclient"
	"context"
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

func makeEndpoint(apiBaseURL string) (string, error) {
	address, err := url.Parse(apiBaseURL)
	if err != nil || address.Scheme == "" || address.Host == "" {
		return "", fmt.Errorf("bad --apiBaseUrl: %q", apiBaseURL)
	}
	return url.JoinPath(address.String(), apiPath)
}

func buildAPIClient(apiBaseURL string) (*httpclient.APIClient, error) {
	endpoint, err := makeEndpoint(apiBaseURL)
	if err != nil {
		return nil, err
	}
	return httpclient.NewAPIClient(httpclient.WithEndpoint(endpoint)), nil
}

func fetchWordPair(ctx context.Context, ac *httpclient.APIClient) (*httpclient.WordPair, error) {
	wp, err := ac.FetchWordPair(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get word pair: %w", err)
	}
	return wp, nil
}
