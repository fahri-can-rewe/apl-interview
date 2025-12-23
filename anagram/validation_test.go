package anagram

import "testing"

func TestValidateWordPair(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid word pair", []string{"listen", "silent"}, false},
		{"non-alphabetic word pair", []string{"listen1", "silent1"}, true},
		{"different length word pair", []string{"listentome", "silent"}, true},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			err := ValidateWordPair(testCase.args[0], testCase.args[1])
			if (err != nil) != testCase.wantErr {
				t.Errorf("validateWordPair(%v) error = %v, wantErr %v", testCase.args, err, testCase.wantErr)
			}
			if err == nil && testCase.wantErr {
				t.Fatalf("buildAPIClient(%q) expected error, got nil", testCase.args)
			}
		})

	}
}
