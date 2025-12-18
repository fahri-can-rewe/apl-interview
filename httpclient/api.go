package httpclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WordPair struct {
	FirstWord  string `json:"word1"`
	SecondWord string `json:"word2"`
}

type APIClient struct {
	HTTPClient *http.Client
	Endpoint   string
}

func NewAPIClientWithHTTP(client *http.Client, endpoint string) *APIClient {
	return &APIClient{
		HTTPClient: client,
		Endpoint:   endpoint,
	}
}

func (client *APIClient) FetchWordPair() (*WordPair, error) {
	response, err := client.HTTPClient.Get(client.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("error making GET call to %s: %w", client.Endpoint, err)
	}
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("warning: failed to close response body: %v", cerr)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status code: %d", response.StatusCode)
	}

	var wordPair WordPair
	err = json.NewDecoder(response.Body).Decode(&wordPair)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return &wordPair, nil
}
