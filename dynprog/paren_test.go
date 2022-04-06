package dynprog

import "testing"

type matdim struct{ w, h int }

func (a matdim) OpCost(b matdim) float64 {
	return float64(a.w * b.h) // new dims
}

func (a matdim) IfApplyOp(b matdim) matdim {
	return matdim{a.w, b.h}
}

var _ OpData[matdim] = matdim{}

func TestBestParenthesizeCost(t *testing.T) {
	tests := []struct {
		input []matdim
		want  float64
	}{
		{
			input: []matdim{
				{1, 4},
				{4, 3},
				{3, 100},
				{100, 1},
			},
			want: 7,
		},
		{
			input: []matdim{
				{1, 100},
				{100, 1},
				{1, 100},
			},
			want: 101,
		},
	}

	for _, test := range tests {
		got := BestParenthesizeCost(test.input)
		if got != test.want {
			t.Errorf("input %v: got %g want %g", test.input, got, test.want)
		}
	}
}
