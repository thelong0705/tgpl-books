package anagram

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		s1  string
		s2  string
		res bool
	}{
		{
			"abc",
			"cba",
			true,
		},
		{
			"abc",
			"cbad",
			false,
		},
		{
			"abc",
			"caa",
			false,
		},
		{
			"a",
			"a",
			true,
		},
		{
			"",
			"",
			true,
		},
	}

	for _, testCase := range testCases {
		got := IsAnagram(testCase.s1, testCase.s2)
		if got != testCase.res {
			t.Errorf("expected %v got %v", testCase.res, got)
		}
	}
}
