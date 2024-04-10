package shopping

import "math"

// formula:
// g(i, S) = min(C(i,k) + g(k, S - {k}))
//      k is a subset of S

// we first calculate the last leg then ie (k, 1), then move on upwards.
// use bitwise operations to store the value, as to avoid duplication
// mem 2d array would have an size of [n][1<<n]
// where first [] would be pos and second [] would be mask
// mask is basically  a binary number indicating nodes we have visited
// if we haven't visited any nodes then mask would be 00000, and if
// we have visited all the cities then it would be (1<<n)-1 or 11111

func FindShortestPath(graph [][]int) int {
	mem := make([][]int, len(graph))
	for idx := range mem {
		mem[idx] = make([]int, (1 << len(graph)))
	}
	return calculate(1, 0, mem, graph)
}

func calculate(mask, pos int, mem, graph [][]int) int {
	visitedAll := (1 << len(graph)) - 1
	if mask == visitedAll {
		return graph[pos][0]
	}
	if mem[pos][mask] != 0 {
		return mem[pos][mask]
	}
	minDis := math.MaxInt
	for node := 0; node < len(graph); node++ {
		if (mask & (1 << node)) != 0 {
			continue
		}
		dis := calculate(mask|(1<<node), node, mem, graph) + graph[pos][node]
		minDis = min(minDis, dis)
	}
	mem[pos][mask] = minDis
	return minDis
}
