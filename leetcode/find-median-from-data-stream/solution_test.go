package leetcode

import "testing"

func expect(t *testing.T, got, want float64) {
	if want != got {
		t.Helper()
		t.Errorf("got %g want %g", got, want)
	}
}

func TestMedianFinderFailed1(t *testing.T) {
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)
	expect(t, finder.FindMedian(), 1.5)
	finder.AddNum(3)
	expect(t, finder.FindMedian(), 2)
}
