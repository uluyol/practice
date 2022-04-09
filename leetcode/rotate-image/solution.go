package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	k := n / 2
	for ring := 0; ring < k; ring++ {
		topIt := newIter(n, ring, iterTop)
		rightIt := newIter(n, ring, iterRight)
		botIt := newIter(n, ring, iterBot)
		leftIt := newIter(n, ring, iterLeft)
		for topIt.ok {
			t := matrix[topIt.i][topIt.j]
			matrix[topIt.i][topIt.j] = matrix[leftIt.i][leftIt.j]
			matrix[leftIt.i][leftIt.j] = matrix[botIt.i][botIt.j]
			matrix[botIt.i][botIt.j] = matrix[rightIt.i][rightIt.j]
			matrix[rightIt.i][rightIt.j] = t

			topIt.next()
			rightIt.next()
			botIt.next()
			leftIt.next()
		}
	}
}

type iterPos int8

const (
	iterTop iterPos = iota
	iterRight
	iterBot
	iterLeft
)

type iter struct {
	n    int
	ring int
	side iterPos

	i, j int
	ok   bool
}

func (it *iter) next() {
	switch it.side {
	case iterTop:
		it.j++
		if it.j == it.n-it.ring-1 {
			it.ok = false
		}
	case iterRight:
		it.i++
		if it.i == it.n-it.ring-1 {
			it.ok = false
		}
	case iterBot:
		it.j--
		if it.j == it.ring {
			it.ok = false
		}
	case iterLeft:
		it.i--
		if it.i == it.ring {
			it.ok = false
		}
	}
}

func newIter(n, ring int, side iterPos) *iter {
	var i, j int
	switch side {
	case iterTop:
		i, j = ring, ring
	case iterRight:
		i, j = ring, n-ring-1
	case iterBot:
		i, j = n-ring-1, n-ring-1
	case iterLeft:
		i, j = n-ring-1, ring
	}
	return &iter{
		n, ring, side, i, j, true,
	}
}
