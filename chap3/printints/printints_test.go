package printints

import "testing"

func TestIntsToString(t *testing.T) {
	testCases := []struct {
		input  []int
		output string
	}{
		{
			[]int{1, 2, 3},
			"[1, 2, 3]",
		},
		{
			[]int{},
			"[]",
		},
	}

	for _, testCase := range testCases {
		got := intsToString(testCase.input)
		if got != testCase.output {
			t.Errorf("expected %s got %s", got, testCase.output)
		}
	}
}
