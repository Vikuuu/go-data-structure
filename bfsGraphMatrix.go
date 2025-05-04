package datstr

import "slices"

func BFSGraphMatrix(graph [][]int, source, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := NewQueue[int]()
	q.Enqueue(source)

	for q.Len() != 0 {
		curr := q.Deque()
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(graph); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}
		seen[curr] = true
	}

	// build it backwards
	//
	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) != 0 {
		out = append(out, source)
		slices.Reverse(out)
		return out
	}
	return out
}
