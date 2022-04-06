package dynprog

import (
	"reflect"
	"testing"
)

func TestSplitIntoLines(t *testing.T) {
	tests := []struct {
		text      string
		lineWidth int
		lines     []string
	}{
		{
			text:      "",
			lineWidth: 0,
			lines:     nil,
		},
		{
			text:      "abc def  ghi  klh js tta",
			lineWidth: 6,
			lines:     []string{"abc", "def", "ghi", "klh", "js tta"},
		},
		{
			text:      "abc def  ghi  klh js tta",
			lineWidth: 7,
			lines:     []string{"abc def", "ghi klh", "js tta"},
		},
		{
			text:      "a be     cdef zzz 1234567",
			lineWidth: 9,
			lines:     []string{"a be", "cdef zzz", "1234567"},
		},
	}

	for _, test := range tests {
		lines := SplitIntoLines(test.text, test.lineWidth)
		if len(lines) == 0 && len(test.lines) == 0 {
			continue
		}
		if !reflect.DeepEqual(lines, test.lines) {
			t.Errorf("text: %q, lineWidth: %d, want: %#v got: %#v",
				test.text, test.lineWidth, test.lines, lines)
		}
	}
}
