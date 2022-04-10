package leetcode

func pacificAtlantic(heights [][]int) [][]int {
	G := constructGraph(heights)
	pacificID := NodeID(len(G.Nodes) - 2)
	atlanticID := NodeID(len(G.Nodes) - 1)
	flowToPacific := dfs(G, pacificID)
	flowToAtlantic := dfs(G, atlanticID)
	var coordFlowToBoth [][]int
	for i := range flowToPacific[:len(G.Nodes)-2] {
		if flowToPacific[i] && flowToAtlantic[i] {
			coordFlowToBoth = append(coordFlowToBoth, []int{G.Nodes[i].X, G.Nodes[i].Y})
		}
	}
	return coordFlowToBoth
}

func dfs(G *Graph, from NodeID) []bool {
	touched := make([]bool, len(G.Nodes))
	dfsImpl(G, from, touched)
	return touched
}

func dfsImpl(G *Graph, from NodeID, touched []bool) {
	touched[from] = true
	for _, to := range G.AdjTo[from] {
		if touched[to] {
			continue
		}
		touched[to] = true
		dfsImpl(G, to, touched)
	}
}

func constructGraph(heights [][]int) *Graph {
	// Create a directed graph where edges flow from
	// LOWER height to HIGHER height.
	//
	// The purpose is the allow a DFS from each ocean
	// to reach the peaks.

	G := &Graph{
		Nodes: make([]Coord, len(heights)*len(heights[0])+2),
		AdjTo: make([][]NodeID, len(heights)*len(heights[0])+2),
	}

	for i := range heights {
		row := i * len(heights[0])
		for j := range heights[0] {
			G.Nodes[row+j] = Coord{i, j}

			tryAddEdgeTo := func(x, y int) {
				//                 if x < 0 || (y < 0 && x < len(heights)-1){
				//                     // Pacific
				//                 }
				if x < 0 || y < 0 || x >= len(heights) || y >= len(heights[0]) {
					return
				}
				if heights[i][j] <= heights[x][y] {
					// Add edge
					G.AdjTo[row+j] = append(G.AdjTo[row+j], NodeID(x*len(heights[0])+y))
				}
			}

			tryAddEdgeTo(i-1, j)
			tryAddEdgeTo(i+1, j)
			tryAddEdgeTo(i, j-1)
			tryAddEdgeTo(i, j+1)
		}
	}

	pacificID := NodeID(len(G.Nodes) - 2)
	atlanticID := NodeID(len(G.Nodes) - 1)

	G.Nodes[len(G.Nodes)-2] = Coord{-1, -1} // Pacific
	G.Nodes[len(G.Nodes)-1] = Coord{-2, -2} // Atlantic

	endJ := len(heights[0]) - 1
	for i := range heights {
		G.AdjTo[pacificID] = append(G.AdjTo[pacificID], NodeID(i*len(heights[0])))
		G.AdjTo[atlanticID] = append(G.AdjTo[atlanticID], NodeID(i*len(heights[0])+endJ))
	}
	endRow := (len(heights) - 1) * len(heights[0])
	for j := range heights[0] {
		G.AdjTo[pacificID] = append(G.AdjTo[pacificID], NodeID(j))
		G.AdjTo[atlanticID] = append(G.AdjTo[atlanticID], NodeID(endRow+j))
	}
	return G
}

type Coord struct {
	X, Y int
}

type Graph struct {
	Nodes []Coord
	AdjTo [][]NodeID
}

type NodeID int
