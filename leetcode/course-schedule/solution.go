package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjTo := make([][]int, numCourses)
	for _, edge := range prerequisites {
		adjTo[edge[0]] = append(adjTo[edge[0]], edge[1])
	}

	visited := make([]color, len(adjTo))
	var q []int
	for startCourse := range adjTo {
		if visited[startCourse] == black {
			continue
		}
		q = append(q, startCourse)
		for len(q) != 0 {
			node := q[len(q)-1]
			// We only pop after recursing to all children
			// At pop time, mark black
			if visited[node] == gray {
				visited[node] = black
				q = q[:len(q)-1]
				continue
			}
			visited[node] = gray
			for _, dst := range adjTo[node] {
				switch visited[dst] {
				case gray:
					return false
				case black:
					continue
				}
				q = append(q, dst)
			}
		}
	}
	return true
}

type color uint8

const (
	white color = iota
	gray
	black
)
