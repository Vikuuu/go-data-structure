package datstr

import (
	"math"
	"slices"
)

func hasUnvisited(seen []bool, dists []int) bool {
	has := false

	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < math.MaxInt {
			has = true
		}
	}

	return has
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDist := math.MaxInt

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}
		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}
	return idx
}

func DijkstraList(
	source, sink int,
	arr WeightedAdjacencyList,
) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	for i := range prev {
		prev[i] = -1
	}
	dists := make([]int, len(arr))
	for i := range dists {
		dists[i] = math.MaxInt
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge["to"]] {
				continue
			}
			dist := dists[curr] + edge["weight"]
			if dist < dists[edge["to"]] {
				dists[edge["to"]] = dist
				prev[edge["to"]] = curr
			}
		}
	}

	out := []int{}
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source)
	slices.Reverse(out)

	return out
}
