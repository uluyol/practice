package dynprog

import (
	"strings"
)

type lcsBest struct {
	n            int
	nexti, nextj int
	c            byte
	hasC         bool
	ok           bool
}

type lcsFinder struct {
	a, b string
	t    []lcsBest
}

func (f *lcsFinder) get(i, j int) *lcsBest {
	return &f.t[len(f.a)*j+i]
}

func (f *lcsFinder) find(i, j int) (n int) {
	// defer func() {
	// 	fmt.Println(i, j, n)
	// }()
	if i >= len(f.a) || j >= len(f.b) {
		n = 0
		return n
	}

	if b := f.get(i, j); b.ok {
		n = b.n
		return n
	}

	if f.a[i] == f.b[j] {
		n = 1 + f.find(i+1, j+1)
		*f.get(i, j) = lcsBest{
			n:     n,
			nexti: i + 1,
			nextj: j + 1,
			c:     f.a[i],
			hasC:  true,
			ok:    true,
		}
		return n
	}
	n1 := f.find(i+1, j)
	n2 := f.find(i, j+1)
	if n1 >= n2 {
		*f.get(i, j) = lcsBest{
			n:     n1,
			nexti: i + 1,
			nextj: j,
			ok:    true,
		}
		n = n1
	} else {
		*f.get(i, j) = lcsBest{
			n:     n2,
			nexti: i,
			nextj: j + 1,
			ok:    true,
		}
		n = n2
	}
	return n
}

func LongestCommonSubsequence(a, b string) string {
	if a == "" || b == "" {
		return ""
	}
	f := lcsFinder{
		a: a,
		b: b,
		t: make([]lcsBest, len(a)*len(b)),
	}
	f.find(0, 0)
	var sb strings.Builder
	var i, j int
	for {
		if i == len(a) || j == len(b) {
			break
		}
		best := f.get(i, j)
		if !best.ok {
			panic("invalid state: !best.ok")
		}
		if best.hasC {
			sb.WriteByte(best.c)
		}
		i = best.nexti
		j = best.nextj
	}
	return sb.String()
}
