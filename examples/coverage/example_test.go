package main

import "testing"

func TestSize(t *testing.T) {
	// test cases
	var tests = []struct {
		in  int
		out string
	}{
		{-1, "negative"},
		{0, "zero"},
	}

	// run each test case
	for _, tt := range tests {
		res := size(tt.in)
		if res != tt.out {
			t.Errorf("size(%d) = %s; want %s", tt.in, res, tt.out)
		}
	}
}
