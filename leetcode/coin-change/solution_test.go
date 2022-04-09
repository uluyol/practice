package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			[]int{1, 2, 5}, 11, 3,
		},
		{
			[]int{2}, 3, -1,
		},
	}

	for _, test := range tests {
		t.Logf("amount = %d coins = %v", test.amount, test.coins)
		got := coinChange(test.coins, test.amount)
		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}
