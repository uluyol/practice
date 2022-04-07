package dynprog

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{"", "", ""},
		{"a", "", ""},
		{"", "b", ""},
		{"a", "a", "a"},
		{"a", "A", ""},
		{"aA", "A", "A"},
		{"aA", "a", "a"},
		{"abcdef", "cdf", "cdf"},
		{"abcdef", "zcydcf1", "cdf"},
	}

	for _, test := range tests {
		t.Logf("a = %q", test.a)
		t.Logf("b = %q", test.b)
		got := LongestCommonSubsequence(test.a, test.b)
		if got != test.want {
			t.Errorf("got %q want %q", got, test.want)
		}
	}
}
