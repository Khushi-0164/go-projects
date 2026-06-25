package main

import "testing"

func TestValidURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"https://google.com", true},
		{"https://github.com", true},
		{"google.com", false},
	}

	for _, test := range tests {

		got := IsValidURL(test.url)

		if got != test.expected {
			t.Errorf(
				"expected %v, got %v",
				test.expected,
				got,
			)
		}
	}
}
