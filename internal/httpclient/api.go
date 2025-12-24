package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WordPair struct {
	FirstWord  string `json:"word1"`
	SecondWord string `json:"word2"`
}

type WordPairFetcher interface {
	FetchWordPair(ctx context.Context) (*WordPair, error)
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type APIClient struct {
	doer     Doer
	endpoint string
}

type Option func(*APIClient)

func WithEndpoint(ep string) Option         { return func(c *APIClient) { c.endpoint = ep } }
func WithHTTPClient(hc *http.Client) Option { return func(c *APIClient) { c.doer = hc } }
func WithTimeout(d time.Duration) Option {
	return func(client *APIClient) {
		if hc, isOk := client.doer.(*http.Client); isOk {
			hc.Timeout = d
		}
	}
}

func NewAPIClient(options ...Option) *APIClient {
	client := &APIClient{
		doer: &http.Client{Timeout: 5 * time.Second},
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (client *APIClient) FetchWordPair(ctx context.Context) (*WordPair, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := client.doer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", client.endpoint, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var wp WordPair
	if err := json.NewDecoder(resp.Body).Decode(&wp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &wp, nil
}
