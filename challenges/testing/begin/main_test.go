// challenges/testing/begin/main_test.go
package main

import (
	"testing"
)

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/", want: 1},
		"two":   {input: "a2 3c2 ^ &/", want: 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("Test failed for input: %v. Expected: %v | Received: %v", tc.input, tc.want, got)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/", want: 3},
		"two":   {input: "a2 3c2657 ^ &/", want: 6},
	}

	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := nc.count(tc.input); got != tc.want {
				t.Errorf("Test failed for input: %v. Expected: %v | Received: %v", tc.input, tc.want, got)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/", want: 6},
		"two":   {input: "a2!@!! 3c2657 ^ &/", want: 10},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := sc.count(tc.input); got != tc.want {
				t.Errorf("Test failed for input: %v. Expected: %v | Received: %v", tc.input, tc.want, got)
			}
		})
	}
}
