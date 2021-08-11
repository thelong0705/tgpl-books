package remove_adjacent_duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveAdjacentDuplicates(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "duplicated in middle",
			input:  []string{"test", "test1", "test1", "test2"},
			expect: []string{"test", "test1", "test2"},
		},
		{
			name:   "duplicated in beginning",
			input:  []string{"test", "test", "test1", "test2"},
			expect: []string{"test", "test1", "test2"},
		},
		{
			name:   "duplicated in ending",
			input:  []string{"test", "test1", "test2", "test2"},
			expect: []string{"test", "test1", "test2"},
		},
		{
			name:   "duplicated all",
			input:  []string{"test", "test", "test", "test"},
			expect: []string{"test"},
		},
		{
			name:   "empty slice",
			input:  []string{},
			expect: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := RemoveAdjacentDuplicates(testCase.input)
			if !reflect.DeepEqual(testCase.expect, got) {
				t.Errorf("expected %v got %v", testCase.expect, got)
			}
		})
	}
}
