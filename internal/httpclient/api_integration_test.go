package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const testHTTPTimeout = 100 * time.Millisecond

func TestFetchWordPair_Success(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, _ *http.Request) {
		respWriter.Header().Set("Content-Type", "application/json")
		respWriter.WriteHeader(http.StatusOK)
		_, err := respWriter.Write([]byte(`{"word1":"listen","word2":"silent"}`))
		if err != nil {
			return
		}
	}))
	defer server.Close()
	client := NewAPIClient(
		WithHTTPClient(&http.Client{Timeout: testHTTPTimeout}),
		WithEndpoint(server.URL),
	)

	wordPair, err := client.FetchWordPair(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := &WordPair{
		FirstWord:  "listen",
		SecondWord: "silent",
	}
	if *wordPair != *want {
		t.Fatalf("want %+v, got %+v", want, wordPair)
	}
}

func TestFetchWordPair_BuildRequestError(t *testing.T) {
	incorrectURL := "http://[::1]:namedport"
	client := NewAPIClient(WithEndpoint(incorrectURL))

	_, err := client.FetchWordPair(context.Background())

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "build request:") {
		t.Fatalf("error = %q, want it to contain %q", err, "build request:")
	}
}

func TestFetchWordPair_ServerError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		respWriter.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClient(WithHTTPClient(httpClient), WithEndpoint(server.URL))

	wordPair, err := client.FetchWordPair(context.Background())

	if err == nil {
		t.Fatal("want error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_InvalidJSON(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		respWriter.WriteHeader(http.StatusOK)
		_, err := respWriter.Write([]byte(`{"word1": " brokenJSON",`))
		if err != nil {
			return
		}
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClient(WithHTTPClient(httpClient), WithEndpoint(server.URL))

	wordPair, err := client.FetchWordPair(context.Background())

	if err == nil {
		t.Fatal("want JSON decode error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_HTTPError(t *testing.T) {
	t.Parallel()
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClient(WithHTTPClient(httpClient), WithEndpoint("http://127.0.0.1:0"))

	wordPair, err := client.FetchWordPair(context.Background())

	if err == nil {
		t.Fatal("want HTTP error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_Timeout(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		time.Sleep(200 * time.Millisecond)
		respWriter.WriteHeader(http.StatusOK)
		_, err := respWriter.Write([]byte(`{"word1":"a","word2":"b"}`))
		if err != nil {
			return
		}
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: 50 * time.Millisecond}
	client := NewAPIClient(WithHTTPClient(httpClient), WithEndpoint(server.URL), WithTimeout(100*time.Millisecond))

	wordPair, err := client.FetchWordPair(context.Background())

	if err == nil {
		t.Fatal("want timeout error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}
