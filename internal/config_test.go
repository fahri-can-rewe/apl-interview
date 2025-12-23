package internal

import (
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
			conf, err := ParseArgs(testCase.args)
			if (err != nil) != testCase.wantErr {
				t.Fatalf("parseArgs(%v) error = %v, wantErr %v", testCase.args, err, testCase.wantErr)
			}
			if conf.APIBaseURL != testCase.want {
				t.Errorf("parseArgs(%v) = %v, want %v", testCase.args, conf.APIBaseURL, testCase.want)
			}
		})
	}
}

func TestMakeEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		wantURL     string
		expectError bool
		errorMsg    string
	}{
		{
			name:     "basic_url",
			inputURL: "https://api.test",
			wantURL:  "https://api.test/word-pair",
		},
		{
			name:     "url_with_trailing_slash",
			inputURL: "https://www.api.test/",
			wantURL:  "https://www.api.test/word-pair",
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

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := MakeEndpoint(tc.inputURL)
			if tc.expectError {
				if err == nil {
					t.Fatalf("makeEndpoint(%q) expected error, got nil", tc.inputURL)
				}
				if tc.errorMsg != "" && !strings.Contains(err.Error(), tc.errorMsg) {
					t.Fatalf("makeEndpoint(%q) error = %v, want containing %q", tc.inputURL, err, tc.errorMsg)
				}
				return
			}
			if err != nil {
				t.Fatalf("makeEndpoint(%q) unexpected error: %v", tc.inputURL, err)
			}
			if got != tc.wantURL {
				t.Fatalf("makeEndpoint(%q) = %q, want %q", tc.inputURL, got, tc.wantURL)
			}
		})
	}
}
