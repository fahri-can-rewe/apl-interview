package main

import (
	"apl-interview/httpclient"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchWordPair_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/word-pair" {
			t.Fatalf("path=%q", req.URL.Path)
		}
		respWriter.Header().Set("Content-Type", "application/json")
		respWriter.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(respWriter, `{"word1":"listen","word2":"silent"}`)
		if err != nil {
			return
		}
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: time.Second}
	apiClient := httpclient.NewAPIClientWithHTTP(httpClient, server.URL+"/word-pair")

	wp, err := fetchWordPair(apiClient)

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if wp.FirstWord != "listen" || wp.SecondWord != "silent" {
		t.Fatalf("unexpected payload: %+v", wp)
	}
}

func TestFetchWordPair_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/word-pair" {
			t.Fatalf("path=%q", req.URL.Path)
		}
		respWriter.Header().Set("Content-Type", "application/json")
		respWriter.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	httpClient := &http.Client{Timeout: time.Second}
	apiClient := httpclient.NewAPIClientWithHTTP(httpClient, server.URL+"/word-pair")

	_, err := fetchWordPair(apiClient)

	if err == nil {
		t.Fatalf("failed to get word pair: %v", err)
	}
}
