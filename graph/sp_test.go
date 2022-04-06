package graph

import (
	"reflect"
	"testing"
)

type srcAndPaths struct {
	src   int
	paths [][]int
}

type spTest struct {
	name    string
	graph   *G[int]
	samples []srcAndPaths
}

type spFunc func(int, []map[int]float64) Preds

func bellmanFordNoNegCycles(src int, edges []map[int]float64) Preds {
	preds, _ := BellmanFordShortestPaths(src, edges)
	return preds
}

func (st *spTest) Run(t *testing.T, spFunc spFunc) {
	for _, sample := range st.samples {
		preds := ShortestPaths(sample.src, st.graph.Edges)
		for _, path := range sample.paths {
			dst := path[len(path)-1]
			got := preds.PathTo(dst)
			if len(path) == 0 && len(got) == 0 {
				continue
			}
			if !reflect.DeepEqual(got, path) {
				t.Errorf("%d->%d: got %v want %v", sample.src, dst, got, path)
			}
		}
	}
}

func testNonNegativeSP(t *testing.T, spFunc spFunc) {
	t.Helper()
	subTests := []spTest{
		{
			name: "Easy",
			graph: New([]int{0, 1, 2, 3, 4}, []map[int]float64{
				0: {
					1: 2,
					2: 1,
				},
				1: {
					0: 1,
					2: 5,
				},
				2: {
					3: 1,
					4: 5,
				},
				3: {
					4: 1,
				},
				4: {},
			}),
			samples: []srcAndPaths{
				{
					src: 0,
					paths: [][]int{
						{0, 2, 3, 4},
						{0, 2},
						{0, 1},
					},
				},
				{
					src: 1,
					paths: [][]int{
						{1, 0, 2, 3, 4},
						{1, 0, 2},
					},
				},
			},
		},
	}

	for _, st := range subTests {
		t.Run(st.name, func(t *testing.T) { st.Run(t, spFunc) })
	}
}

func TestShortestPaths(t *testing.T) {
	testNonNegativeSP(t, ShortestPaths)
}

func TestBellmanFordShortestPathsNoNegCycles(t *testing.T) {
	testNonNegativeSP(t, bellmanFordNoNegCycles)
}
