package leetcode

func uniquePaths(m int, n int) int {
	dim := m + 1
	if n > m {
		dim = n + 1
	}
	s := state{cache: make([]int, dim*dim), dim: dim}
	return s.solve(m, n)
}

type state struct {
	cache []int
	dim   int
}

func (s *state) get(i, j int) *int {
	return &s.cache[i*s.dim+j]
}

func (s *state) solve(m int, n int) int {
	if n < m {
		m, n = n, m
	}
	// fmt.Println(m, n)
	if m == 1 || n == 1 {
		return 1
	}
	if m <= 0 || n <= 0 {
		return 0
	}
	slot := s.get(m, n)
	if *slot != 0 {
		// fmt.Println(m, n, "cached")
		return *slot
	}
	// defer func() {
	// 	fmt.Println(m, n, *slot)
	// }()
	*slot = s.solve(m, n-1) + s.solve(m-1, n)
	return *slot
}
