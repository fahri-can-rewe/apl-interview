package internal

import (
	"apl-interview/httpclient"
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestFetchWordPair_Success(t *testing.T) {
	body := bytes.NewBufferString(`{"word1":"listen","word2":"silent"}`)
	apiClient := generateAPIClient(http.StatusOK, io.NopCloser(body))

	wp, err := fetchWordPair(context.Background(), apiClient)

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if wp.FirstWord != "listen" || wp.SecondWord != "silent" {
		t.Fatalf("unexpected payload: %+v", wp)
	}
}

func TestFetchWordPair_Failure(t *testing.T) {
	apiClient := generateAPIClient(http.StatusInternalServerError, io.NopCloser(bytes.NewBuffer(nil)))

	_, err := fetchWordPair(context.Background(), apiClient)

	if err == nil {
		t.Fatalf("failed to get word pair: %v", err)
	}
}

func generateAPIClient(statusCode int, body io.ReadCloser) *httpclient.APIClient {
	rt := stubRoundTrip(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: statusCode,
			Body:       body,
			Header:     make(http.Header),
			Request:    req,
		}, nil
	})
	httpClient := &http.Client{Timeout: time.Second, Transport: rt}
	apiClient := httpclient.NewAPIClient(
		httpclient.WithHTTPClient(httpClient),
		httpclient.WithEndpoint("https://example.com/word-pair"),
	)
	return apiClient
}

type stubRoundTrip func(*http.Request) (*http.Response, error)

func (f stubRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }
