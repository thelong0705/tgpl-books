package reverse

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		input [6]int
		want  [6]int
	}{
		{
			input: [6]int{1, 2, 3, 4, 5, 6},
			want:  [6]int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		reverse(&testCase.input)
		got := testCase.input
		if got != testCase.want {
			t.Errorf("Expected %v got %v", testCase.want, got)
		}
	}
}
