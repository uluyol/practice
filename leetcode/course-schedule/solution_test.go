package leetcode

import "testing"

func TestCanFinish_BasicCycle(t *testing.T) {
	got := canFinish(2, [][]int{
		{0, 1},
		{1, 0},
	})
	if got != false {
		t.Errorf("have a cycle but told that we can finish")
	}
}
