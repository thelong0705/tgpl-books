package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name                string
		inputSlice          []int
		inputRotateElements int
		wantSlice           []int
	}{
		{
			name:                "rotate by 0 elements",
			inputSlice:          []int{1, 2, 3, 4, 5},
			inputRotateElements: 0,
			wantSlice:           []int{1, 2, 3, 4, 5},
		},
		{
			name:                "rotate by number of elements that smaller than len",
			inputSlice:          []int{1, 2, 3, 4, 5},
			inputRotateElements: 2,
			wantSlice:           []int{3, 4, 5, 1, 2},
		},

		{
			name:                "rotate by number of elements that equal to len",
			inputSlice:          []int{1, 2, 3, 4, 5},
			inputRotateElements: 5,
			wantSlice:           []int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := Rotate(testCase.inputSlice, testCase.inputRotateElements)
			if !reflect.DeepEqual(testCase.wantSlice, got) {
				t.Errorf("want %v got %v", testCase.wantSlice, got)
			}
		})
	}
}
