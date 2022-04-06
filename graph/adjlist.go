package graph

type G[T any] struct {
	Nodes []T
	Edges []map[int]float64
}

func New[T any](nodes []T, edges []map[int]float64) *G[T] {
	return &G[T]{
		Nodes: nodes,
		Edges: edges,
	}
}
