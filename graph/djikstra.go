package graph

import (
	"github.com/uluyol/practice/basicds"
)

type Preds []int

func (p Preds) walkPredsOf(dst int, fn func(cur int)) {
	for p[dst] != dst {
		fn(dst)
		dst = p[dst]
	}
	fn(dst)
}

func (p Preds) PathTo(dst int) []int {
	size := 0
	p.walkPredsOf(dst, func(cur int) { size++ })
	path := make([]int, size)
	pos := size - 1
	p.walkPredsOf(dst, func(cur int) {
		path[pos] = cur
		pos--
	})
	return path
}

func ShortestPaths(src int, edges []map[int]float64) Preds {
	type weightedTo struct {
		dst    int
		pred   int
		weight float64
	}
	workq := basicds.Heap[weightedTo]{
		Less: func(a, b weightedTo) bool {
			if a.weight == b.weight {
				return a.dst > b.dst
			}
			return a.weight > b.weight // min priority queue
		},
	}
	// Not quite right. Really want update-key for workq
	visited := make([]bool, len(edges))
	workq.Push(weightedTo{src, src, 0})
	preds := make(Preds, len(edges))
	for workq.Len() > 0 {
		cur := workq.PopMax() // pop min, see Less func used
		if visited[cur.dst] {
			continue
		}
		preds[cur.dst] = cur.pred
		visited[cur.dst] = true

		for next, weight := range edges[cur.dst] {
			if visited[next] {
				continue
			}
			workq.Push(weightedTo{next, cur.dst, cur.weight + weight})
		}
	}

	return preds
}
