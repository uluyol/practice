package leetcode

import "testing"

func TestLengthLongestPath(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
	}

	for _, test := range tests {
		got := lengthLongestPath(test.input)
		if got != test.want {
			t.Errorf("input = %s\ngot %d want %d", test.input,
				got, test.want)
		}
	}
}
