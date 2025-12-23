package main

import (
	"apl-interview/anagram"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{"default", nil, defaultAPIBaseURL, false},
		{"custom", []string{"--apiBaseUrl", "https://example.com/base"}, "https://example.com/base", false},
		{"unknown", []string{"--nope"}, "", true},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			conf, err := parseArgs(testCase.args)
			if (err != nil) != testCase.wantErr {
				t.Fatalf("parseArgs(%v) error = %v, wantErr %v", testCase.args, err, testCase.wantErr)
			}
			if conf.APIBaseURL != testCase.want {
				t.Errorf("parseArgs(%v) = %v, want %v", testCase.args, conf.APIBaseURL, testCase.want)
			}
		})
	}
}

func TestBuildAPIClient(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		wantURL     string
		expectError bool
		errorMsg    string // optional expected error message substring
	}{
		{
			name:     "basic_url",
			inputURL: "https://api.test",
			wantURL:  "https://api.test/word-pair",
		},
		{
			name:     "url_with_trailing_slash",
			inputURL: "https://api.test/",
			wantURL:  "https://api.test/word-pair",
		},
		{
			name:     "url_with_base_path",
			inputURL: "https://api.test/base",
			wantURL:  "https://api.test/base/word-pair",
		},
		{
			name:     "url_with_base_path_and_slash",
			inputURL: "https://api.test/base/",
			wantURL:  "https://api.test/base/word-pair",
		},
		{
			name:     "url_with_sub_paths",
			inputURL: "https://api.test/v1/api",
			wantURL:  "https://api.test/v1/api/word-pair",
		},
		{
			name:        "no_scheme",
			inputURL:    "api.test",
			expectError: true,
			errorMsg:    "bad --apiBaseUrl",
		},
		{
			name:        "empty_scheme",
			inputURL:    "://api.test",
			expectError: true,
			errorMsg:    "bad --apiBaseUrl",
		},
		{
			name:        "no_host",
			inputURL:    "https://",
			expectError: true,
			errorMsg:    "bad --apiBaseUrl",
		},
		{
			name:        "empty_string",
			inputURL:    "",
			expectError: true,
			errorMsg:    "bad --apiBaseUrl",
		},
		{
			name:        "invalid_url",
			inputURL:    "://",
			expectError: true,
			errorMsg:    "bad --apiBaseUrl",
		},
		{
			name:        "malformed_url",
			inputURL:    "http://[::1]:namedport",
			expectError: true,
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			_, err := buildAPIClient(testCase.inputURL)
			if testCase.expectError {
				if err == nil {
					t.Errorf("buildAPIClient(%q) expected error, got nil", testCase.inputURL)
				}
				if testCase.errorMsg != "" && !strings.Contains(err.Error(), testCase.errorMsg) {
					t.Errorf("buildAPIClient(%q) error = %v, want containing %q",
						testCase.inputURL, err, testCase.errorMsg)
				}
				return
			}
			if err != nil {
				t.Fatalf("buildAPIClient(%q) unexpected error: %v", testCase.inputURL, err)
			}
		})
	}
}

func TestAreAnagrams(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid word pair", []string{"listen", "silent"}, false},
		{"non-alphabetic word pair", []string{"domain", "amodin"}, false},
		{"different length word pair", []string{"listentome", "silent"}, true},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			result := areAnagrams(testCase.args[0], testCase.args[1], anagram.FreqMapChecker{})
			if result == testCase.wantErr {
				t.Fatalf("areAnagrams(%v) = %v, want %v", testCase.args, result, testCase.wantErr)
			}
		})

	}
}
