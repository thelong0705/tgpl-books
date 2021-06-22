package comma

import "testing"

type testCase struct {
	input  string
	output string
}

func TestComma(t *testing.T) {
	testCases := []testCase{
		{
			"12345",
			"12,345",
		},
		{
			"123456",
			"123,456",
		},
		{
			"1234567",
			"1,234,567",
		},
		{
			"1234",
			"1,234",
		},
		{
			"123",
			"123",
		},
	}

	for _, testCase := range testCases {
		got := commaWithByteBuffer(testCase.input)

		if got != testCase.output {
			t.Errorf("expected %s got %s", testCase.output, got)
		}
	}
}
