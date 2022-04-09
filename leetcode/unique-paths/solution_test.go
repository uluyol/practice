package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{7, 3, 28},
		{10, 10, 48620},
	}

	for _, test := range tests {
		got := uniquePaths(test.m, test.n)
		if got != test.want {
			t.Errorf("m=%d n=%d got %d want %d", test.m, test.n, got, test.want)
		}
	}
}
