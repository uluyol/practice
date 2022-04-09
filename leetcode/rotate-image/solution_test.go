package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, test := range tests {
		t.Logf("input = %v", test.input)
		rotate(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("got:\n%v\nwant:\n%v", test.input, test.want)
		}
	}
}
