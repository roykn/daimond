package main

import (
	"testing"
)

type testpair struct {
	letter string
	line   string
}

var tests = []testpair{
	{"A", "A"},
	{"B", "B B"},
	{"C", "C   C"},
	{"D", "D     D"},
}

func TestGenerateLine(t *testing.T) {
	for _, test := range tests {
		line := generateLine(test.letter)
		if line != test.line {
			t.Error(
				"For", test.letter,
				"expected", test.line,
				"got", line,
			)
		}
	}
}
