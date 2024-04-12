package route

import (
	"container/heap"
	"strconv"
	"strings"
)

// also utilize dijkstra algorithm for this question
// convert the data in the input file into a hash map, this would allow us to easily process it

// various functions required for heap data structure for dijkstra's algorithm

type beacon struct {
	name     string
	distance int
	index    int
}

type beaconHeap []*beacon

// implement heap.Interface

func (b beaconHeap) Len() int {
	return len(b)
}

func (b beaconHeap) Less(i, j int) bool {
	return b[i].distance < b[j].distance
}

func (b beaconHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
	b[i].index, b[j].index = i, j
}

func (b *beaconHeap) Push(x any) {
	itemIndex := len(*b)
	item := x.(*beacon)
	item.index = itemIndex
	*b = append(*b, item)
}

func (b *beaconHeap) Pop() any {
	n := len(*b) - 1
	toReturn := (*b)[n]
	(*b)[n] = nil
	toReturn.index = -1
	*b = (*b)[:n]
	return toReturn
}

func (b *beaconHeap) update(item *beacon, distance int) {
	item.distance = distance
	heap.Fix(b, item.index)
}

func FindShortestPath(data map[string]string, start, end string, stopTime int) (int, error) {
	visited := make(map[string]struct{})
	nameToBeacon := make(map[string]*beacon)
	queue := new(beaconHeap)

	startBeacon := &beacon{name: start, distance: 0}
	nameToBeacon[startBeacon.name] = startBeacon
	heap.Push(queue, startBeacon)

	for queue.Len() > 0 {
		root := heap.Pop(queue).(*beacon)
		if root.name == end {
			return root.distance, nil
		}
		stations := strings.Fields(data[root.name])
		for _, station := range stations {
			sData := strings.Split(station, ":")
			if _, ok := visited[sData[0]]; ok {
				continue
			}
			dist, err := strconv.Atoi(sData[1])
			if err != nil {
				return 0, err
			}
			dist += root.distance

			if root.name != start {
				dist += stopTime
			}
			if addr, ok := nameToBeacon[sData[0]]; ok {
				queue.update(addr, min(addr.distance, dist))
				continue
			}
			b := &beacon{name: sData[0], distance: dist}
			nameToBeacon[sData[0]] = b
			heap.Push(queue, b)
		}
		visited[root.name] = struct{}{}
	}
	// in case we don't find a path to the end node
	return -1, nil
}
