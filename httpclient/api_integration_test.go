package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const testHTTPTimeout = 2 * time.Second

func TestFetchWordPair_Success(t *testing.T) {
	t.Parallel()

	// Assign
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, _ *http.Request) {
		respWriter.Header().Set("Content-Type", "application/json")
		respWriter.WriteHeader(http.StatusOK)
		_, err := respWriter.Write([]byte(`{"word1":"listen","word2":"silent"}`))
		if err != nil {
			return
		}
	}))
	defer server.Close()
	client := NewAPIClientWithHTTP(
		&http.Client{Timeout: testHTTPTimeout},
		server.URL,
	)

	// Act
	wordPair, err := client.FetchWordPair()

	// Assert
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

func TestFetchWordPair_ServerError(t *testing.T) {
	t.Parallel()

	// Assign
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		respWriter.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClientWithHTTP(httpClient, server.URL)

	// Act
	wordPair, err := client.FetchWordPair()

	// Assert
	if err == nil {
		t.Fatal("want error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_InvalidJSON(t *testing.T) {
	t.Parallel()

	// Assign
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		respWriter.WriteHeader(http.StatusOK)
		_, err := respWriter.Write([]byte(`{"word1": " brokenJSON",`))
		if err != nil {
			return
		}
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClientWithHTTP(httpClient, server.URL)

	// Act
	wordPair, err := client.FetchWordPair()

	// Assert
	if err == nil {
		t.Fatal("want JSON decode error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_HTTPError(t *testing.T) {
	t.Parallel()

	// Assign
	httpClient := &http.Client{Timeout: testHTTPTimeout}
	client := NewAPIClientWithHTTP(httpClient, "http://127.0.0.1:0")

	// Act
	wordPair, err := client.FetchWordPair()

	// Assert
	if err == nil {
		t.Fatal("want HTTP error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}

func TestFetchWordPair_Timeout(t *testing.T) {
	t.Parallel()

	// Assign
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
	client := NewAPIClientWithHTTP(httpClient, server.URL)

	// Act
	wordPair, err := client.FetchWordPair()

	// Assert
	if err == nil {
		t.Fatal("want timeout error, got nil")
	}
	if wordPair != nil {
		t.Fatalf("want nil wordPair, got %+v", wordPair)
	}
}
