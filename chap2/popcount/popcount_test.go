package popcount

import (
	"testing"
)

type TestCase struct {
	input uint64
	want  int
}

var testCases []TestCase

func init() {
	testCases = []TestCase{
		{0, 0},
		{1, 1},
		{5, 2},
		{13, 3},
		{113, 4},
		{1013, 8},
	}
}

func TestPopCountLoop(t *testing.T) {
	for _, testCase := range testCases {
		if CountByLoop(testCase.input) != testCase.want {
			t.Errorf("want %d got %d", testCase.input, testCase.want)
		}
	}
}

func TestPopCountTable(t *testing.T) {
	for _, testCase := range testCases {
		if CountByTable(testCase.input) != testCase.want {
			t.Errorf("want %d got %d", testCase.input, testCase.want)
		}
	}
}
