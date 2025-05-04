package datstr

type (
	GraphEdge             map[string]int
	WeightedAdjacencyList [][]GraphEdge
)

func dfsWalk(
	graph WeightedAdjacencyList,
	curr, needle int,
	seen []bool,
	path *[]int,
) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	// recursion
	// pre
	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	// recurse
	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if dfsWalk(graph, edge["to"], needle, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}

func DFSGraphList(graph WeightedAdjacencyList, source, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	dfsWalk(graph, source, needle, seen, &path)

	if len(path) == 0 {
		return []int{}
	}
	return path
}
