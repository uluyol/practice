package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
	}

	for _, test := range tests {
		t.Logf("s = %s", test.s)
		t.Logf("dict = %v", test.dict)
		got := wordBreak(test.s, test.dict)
		if got != test.want {
			t.Errorf("got %t want %t", got, test.want)
		}
	}
}
