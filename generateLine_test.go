package main

import (
	"testing"
)

type testLine struct {
	letter string
	line   string
}

type testOffset struct {
	letter string
	widest string
	offset byte
}

type testGap struct {
	letter string
	gap    int
}

var lines = []testLine{
	{"A", "A"},
	{"B", "B B"},
	{"C", "C   C"},
	{"D", "D     D"},
}

var offsets = []testOffset{
	{"A", "A", 0},
	{"A", "B", 1},
	{"A", "C", 2},
	{"D", "F", 2},
	{"A", "F", 5},
}

var gaps = []testGap{
	{"A", -1},
	{"B", 1},
	{"C", 3},
	{"D", 5},
	{"E", 7},
	{"F", 9},
}

func TestGenerateLine(t *testing.T) {
	for _, test := range lines {
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

func TestOffset(t *testing.T) {
	for _, test := range offsets {
		offset := getOffset(test.letter, test.widest)
		if offset != test.offset {
			t.Error(
				"For", test.letter,
				"with", test.widest, "as widest",
				"expected", test.offset,
				"got", offset,
			)
		}
	}
}

func TestGap(t *testing.T) {
	for _, test := range gaps {
		gap := getGap(test.letter)
		if gap != test.gap {
			t.Error(
				"For", test.letter,
				"expected", test.gap,
				"got", gap,
			)
		}
	}
}
