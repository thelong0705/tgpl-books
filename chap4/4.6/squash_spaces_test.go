package squash_spaces

import (
	"reflect"
	"testing"
)

func TestSquashSpaces(t *testing.T) {
	testCases := []struct {
		name   string
		input  []byte
		output []byte
	}{
		{
			name:   "2 spaces next to each other",
			input:  []byte("A  B C"),
			output: []byte("A B C"),
		},
		{
			name:   "4 spaces next to each other",
			input:  []byte("A    B        C"),
			output: []byte("A B C"),
		},
		{
			name:   "spaces next to tab",
			input:  []byte("A   \tB        C"),
			output: []byte("A B C"),
		},
		{
			name:   "spaces next to japanese spaces",
			input:  []byte("A   　　B        C"),
			output: []byte("A B C"),
		},
		{
			name:   "tab next to tab",
			input:  []byte("A\t\tB        C"),
			output: []byte("A B C"),
		},
		{
			name:   "japanese space next to japanese space",
			input:  []byte("A　　B　　C"),
			output: []byte("A B C"),
		},
		{
			name:   "non ascii code",
			input:  []byte("日    本        語"),
			output: []byte("日 本 語"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := SquashSpaces(testCase.input)

			if !reflect.DeepEqual(testCase.output, got) {
				t.Errorf("want %v got %v", testCase.output, got)
			}
		})
	}
}
