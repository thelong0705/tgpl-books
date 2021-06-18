package basename

import "testing"

type testCase struct {
	input  string
	expect string
}

func TestBaseName(t *testing.T) {
	testCases := []testCase{
		{
			"a/b/c.go",
			"c",
		},
		{
			"c.d.go",
			"c.d",
		},
		{
			"abc",
			"abc",
		},
		{
			"日本/語/テスト.テスト",
			"テスト",
		},
	}

	for _, testCase := range testCases {
		got := Basename(testCase.input)
		if got != testCase.expect {
			t.Errorf("expected %s got %s", testCase.expect, got)
		}
	}

}
