package basicds

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		data []int
		want []int
	}{
		{
			[]int{5, 1, -1, 4},
			[]int{-1, 1, 4, 5},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 0, 3, 2, 5, 4, 7, 6, 9, 8},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		got := append([]int(nil), test.data...)
		HeapSort(got)
		if len(got) == 0 && len(test.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("input %v: got %v want %v", test.data, got, test.want)
		}
	}
}
