package charcount

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestCountRuneAndFreq(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output map[rune]int
	}{
		{
			name:  "count ascii character",
			input: "abc",
			output: map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
		},
		{
			name:  "count ascii character with duplicates",
			input: "aabbbbc",
			output: map[rune]int{
				'a': 2,
				'b': 4,
				'c': 1,
			},
		},
		{
			name:  "count ascii character with spaces",
			input: "a b c",
			output: map[rune]int{
				'a': 1,
				' ': 2,
				'b': 1,
				'c': 1,
			},
		},
		{
			name:  "count unicode characters",
			input: "日本語",
			output: map[rune]int{
				'日': 1,
				'本': 1,
				'語': 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := countRuneFreq(testCase.input)
			if !reflect.DeepEqual(testCase.output, got) {
				t.Errorf("expect %v got %v", testCase.output, got)
			}
		})
	}
}

func TestCountLenFreq(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output [utf8.UTFMax + 1]int
	}{
		{
			name:   "count ascii character",
			input:  "abc",
			output: [utf8.UTFMax + 1]int{0, 3, 0, 0, 0},
		},
		{
			name:   "count unicode character",
			input:  "abc日本語",
			output: [utf8.UTFMax + 1]int{0, 3, 0, 3, 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := countLenFreq(testCase.input)
			if testCase.output != got {
				t.Errorf("expect %v got %v", testCase.output, got)
			}
		})
	}
}
