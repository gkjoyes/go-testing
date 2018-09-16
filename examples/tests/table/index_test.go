package table

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	// test cases
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"ofofof", "f", 1},
	}

	// run each test cases
	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
		}
	}
}
