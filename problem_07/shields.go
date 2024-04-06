package shields

import (
	"cmp"
	"slices"
	"sort"
)

type ColData struct {
	start int
	end   int
}

type inProcess struct {
	data map[ColData]struct{}
}

func newInProcess() *inProcess {
	return &inProcess{data: make(map[ColData]struct{})}
}

func (i *inProcess) add(colStart, colEnd int) {
	i.data[ColData{colStart, colEnd}] = struct{}{}
}

func (i *inProcess) remove(colStart, colEnd int) {
	delete(i.data, ColData{colStart, colEnd})
}

func (i *inProcess) sortedList() []ColData {
	out := make([]ColData, 0, len(i.data))
	for data := range i.data {
		out = append(out, data)
	}
	slices.SortFunc(out, func(a, b ColData) int { return cmp.Compare(a.start, b.start) })
	return out
}

// use line sweep algorithm:
// start from the biggest rectangle to the smallest, this will reduce our complexity
// as we won't need to process every single row, well we could, but we save on that
// get the starting y position, and ending y position, add them both to
// an events array, as we only have two events that can occur
// a rectangle starting from the current lines and a rectangle ending at a current line.
// then we sort the events array, from smallest row to biggest row
// also sort the events array from closing event to opening event if its the same row
// i don't think that matters but we will stick with it for now.
// now we process each event in our events array, and use a new inProcess hashmap
// keep track of the row position for the prev event trigger, can be any event
// on an open event we add colStart and colEnd to the inProcess hashmap
// on a close event we remove colStart and colEnd from the  inProcess hashmap
// now we have a func which calculates the area from each col point
// KEEP TRACK OF PREVIOUS ROW AND COL IN BOTH SWEEPS
//
// (YPos)
//  -------------------------   events = append(events, int[YPos, 1, XPos, XPos+Width])
//  |    					|
//  |    					|
//  |    					|
//  |    					|
//  |    					|
//  |    					|
//  |    					|
//  |    					|
//  -------------------------   events = append(events, int[(YPos + Height), 0, XPos, XPos+Width])
// 							(YPos + Height)

func FindVulnerabilities(shields [][]int, rows, cols int) int {
	var events [][]int
	for sIdx := range shields {
		// rectangle start event
		events = append(events, []int{shields[sIdx][1], 1, shields[sIdx][0], (shields[sIdx][0] + shields[sIdx][2])})
		// rectangle ends event
		events = append(events, []int{(shields[sIdx][1] + shields[sIdx][3]), 0, shields[sIdx][0], (shields[sIdx][0] + shields[sIdx][2])})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	prev := 0
	area := 0
	openIntervals := newInProcess()
	for eIdx := range events {
		curArea := 0
		curPrev := 0
		for _, c := range openIntervals.sortedList() {
			curPrev = max(curPrev, c.start)
			curArea += max(0, (c.end-curPrev)*(events[eIdx][0]-prev))
			curPrev = max(curPrev, c.end)
		}
		if events[eIdx][1] == 0 {
			openIntervals.remove(events[eIdx][2], events[eIdx][3])
		}
		if events[eIdx][1] == 1 {
			openIntervals.add(events[eIdx][2], events[eIdx][3])
		}
		area += curArea
		prev = events[eIdx][0]
	}
	return (rows * cols) - area
}

/*
type HeatShield struct {
	XPos   int // cols
	YPos   int // rows
	Width  int // cols
	Height int // rows
}


func FindVulnerabilities(shields []*HeatShield, rows, cols int) int {
	grid := constructGrid(rows, cols)
	for _, s := range shields {
		for i := s.YPos; i < (s.YPos + s.Height); i++ {
			for j := s.XPos; j < (s.XPos + s.Width); j++ {
				grid[i][j] = true
			}
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !grid[i][j] {
				count++
			}
		}
	}
	return count
}

// constructGrid function constructs a grid of bool type
func constructGrid(rows, cols int) [][]bool {
	grid := make([][]bool, rows)
	for idx := range grid {
		grid[idx] = make([]bool, cols)
	}
	return grid
}
*/
