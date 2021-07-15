package sha256_diff

import (
	"testing"
)

func TestDiff(t *testing.T) {
	testCases := []struct {
		input1       [32]byte
		input2       [32]byte
		expectedDiff int
	}{
		{
			input1:       [32]byte{4},
			input2:       [32]byte{6},
			expectedDiff: 1,
		},
		{
			input1:       [32]byte{4},
			input2:       [32]byte{4},
			expectedDiff: 0,
		},
		{
			input1:       [32]byte{1,2,3},
			input2:       [32]byte{4,5,6},
			expectedDiff: 7,
		},
	}
	for _, testCase := range testCases {
		gotDiff := Diff(testCase.input1, testCase.input2)
		if testCase.expectedDiff != gotDiff {
			t.Errorf("expected %d got %d", testCase.expectedDiff, gotDiff)
		}
	}

}
