package config

import (
	"flag"
	"fmt"
	"io"
	"net/url"
)

const (
	defaultAPIBaseURL = "https://interview.sowula.at"
	apiPath           = "/word-pair"
)

// Config holds CLI configuration derived from flags.
type Config struct {
	APIBaseURL string
}

func ParseArgs(args []string) (Config, error) {
	fs := flag.NewFlagSet("wordpair", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	apiBaseURL := fs.String("apiBaseUrl", defaultAPIBaseURL, "Base URL of the Word-Pair API")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	return Config{APIBaseURL: *apiBaseURL}, nil
}

func MakeEndpoint(apiBaseURL string) (string, error) {
	address, err := url.Parse(apiBaseURL)
	if err != nil {
		return "", fmt.Errorf("bad --apiBaseUrl %q: %w", apiBaseURL, err)
	}
	if address.Scheme == "" || address.Host == "" {
		return "", fmt.Errorf("bad --apiBaseUrl %q: must be absolute (scheme + host)", apiBaseURL)
	}
	return url.JoinPath(address.String(), apiPath)
}
