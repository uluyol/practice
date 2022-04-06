package graph

import "math"

func BellmanFordShortestPaths(src int, edges []map[int]float64) (preds Preds, hasNegativeWeightCycle bool) {
	bestWeight := make([]float64, len(edges))
	for i := range bestWeight {
		bestWeight[i] = math.Inf(1)
	}
	preds = make([]int, len(edges))

	for pass := 0; pass < len(edges)-1; pass++ {
		for v, adj := range edges {
			for v2, w := range adj {
				if t := bestWeight[v] + w; t < bestWeight[v2] {
					bestWeight[v2] = t
					preds[v2] = v
				}
			}
		}
	}

	for v, adj := range edges {
		for v2, w := range adj {
			if t := bestWeight[v] + w; t < bestWeight[v2] {
				hasNegativeWeightCycle = true
				break
			}
		}
	}

	return preds, hasNegativeWeightCycle
}
