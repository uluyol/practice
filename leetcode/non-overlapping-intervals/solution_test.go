package leetcode

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
		debug bool
	}{
		{
			input: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:  2,
		},
		{
			input: [][]int{{1, 2}, {2, 3}},
			want:  0,
		},
		{
			input: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:  1,
		},
		{
			input: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			want:  2,
			debug: false,
		},
		{
			input: [][]int{
				{18, 42}, {-12, -3}, {-83, 66}, {4, 32}, {0, 29}, {62, 72}, {-97, -14}, {24, 87}, {23, 56}, {67, 97}, {14, 48}, {41, 48}, {-59, 74}, {-91, 50}, {35, 97}, {77, 83}, {57, 68}, {-99, 86}, {-27, 16}, {84, 94}, {88, 90}, {91, 93}, {92, 96}, {-78, -24}, {32, 76}, {-90, 7}, {-78, -38}, {-67, 30}, {4, 58}, {35, 36}, {-47, -18}, {-17, -7}, {39, 70}, {85, 86}, {-28, -15}, {91, 97}, {-84, 1}, {30, 71}, {2, 93}, {66, 97}, {94, 97}, {-7, 74}, {-3, 26},
			},
			want:  31,
			debug: false,
		},
	}

	for _, test := range tests {
		// debug = test.debug
		t.Logf("input = %v", test.input)
		got := eraseOverlapIntervals(test.input)
		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}
