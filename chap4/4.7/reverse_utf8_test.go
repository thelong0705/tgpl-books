package reverse_utf8

import (
	"reflect"
	"testing"
)

func TestReverseUTF8(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "ascii string",
			input:    []byte("A B C"),
			expected: []byte("C B A"),
		},
		{
			name:     "non ascii string",
			input:    []byte("日本"),
			expected: []byte("本日"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ReverseUTF8(testCase.input)

			if !reflect.DeepEqual(testCase.expected, testCase.input) {
				t.Errorf("expected %v got %v", testCase.expected, testCase.input)
			}
		})
	}
}
