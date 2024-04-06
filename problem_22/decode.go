package decodepixel

import (
	"io"
	"sort"
)

// use line sweep method, sort the data according to y coord
// add start of the event (start of the rectangle) and end of the event(end of the rectangle) to events array
// create a activeEvents array, it requires use of a custom datatype.
// order of the sweep, first remove, then add then perform operation
// we iterate till previous is different from the current, then we calculate and add to buffer.
// remember to record rectangle with coords in the activeEvents, as each rectangle is needed.

type position struct {
	start, end int
}

type record struct {
	data map[position]int
}

func newRecord() *record {
	return &record{data: make(map[position]int)}
}

func (r *record) add(start, end int) {
	r.data[position{start: start, end: end}] += 1
}

func (r *record) remove(start, end int) {
	if r.data[position{start: start, end: end}] == 1 {
		delete(r.data, position{start: start, end: end})
		return
	}
	r.data[position{start: start, end: end}]--
}

func (r *record) createSubEvents(width int) [][]int {
	var output [][]int
	for key, val := range r.data {
		for i := val; i > 0; i-- {
			output = append(output, []int{1, key.start})
			output = append(output, []int{0, key.end})
		}
	}
	for num := range width {
		output = append(output, []int{2, num})

	}
	sort.Slice(output, func(i, j int) bool {
		if output[i][1] == output[j][1] {
			return output[i][0] < output[j][0]
		}
		return output[i][1] < output[j][1]
	})
	return output
}

func Decode(data [][]int, height, width int, writer io.Writer) {
	// data
	// 0	-	x
	// 1	-	y
	// 2	-	width
	// 3	-	height

	var events [][]int
	for _, c := range data {
		//					heightStart	eventType   start     end
		events = append(events, []int{c[1], 1, c[0], c[0] + c[2]}) // start event
		//					    heightStart	eventType start   end
		events = append(events, []int{c[1] + c[3], 0, c[0], c[0] + c[2]}) // end event
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	prev := 0
	activeEvents := newRecord()
	for _, event := range events {
		if prev != event[0] {
			buffer := make([]byte, 0, width+1)
			accum := 0
			subEvents := activeEvents.createSubEvents(width)
			for _, subEvent := range subEvents {
				switch subEvent[0] {
				case 0:
					accum--
				case 1:
					accum++
				case 2:
					if (accum % 2) == 0 {
						buffer = append(buffer, '.')
						break
					}
					buffer = append(buffer, '#')
				}
			}
			buffer = append(buffer, '\n')
			for i := prev; i < event[0]; i++ {
				writer.Write(buffer)
			}
			prev = event[0]
		}
		if event[1] == 0 {
			activeEvents.remove(event[2], event[3])
			continue
		}
		activeEvents.add(event[2], event[3])
	}
}
