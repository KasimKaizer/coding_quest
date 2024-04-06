package lottery

import "sort"

// maybe use binary search for funnies, we just go through each ticket in the tickets
// matrix then then we go through each ticket and check if its found in winning numbers,
// if so then add 1 to accumulator, at the end of the looping a ticket, we have a  switch statement
// which adds appropriate amount to our total wonAmount.

func CalculateWinnings(tickets [][]int, winNum []int) int {
	sort.Slice(winNum, func(i, j int) bool {
		return winNum[i] < winNum[j]
	})
	totalWinnings := 0
	for _, ticket := range tickets {
		accumulator := 0
		for _, num := range ticket {
			if searchTicket(num, winNum) {
				accumulator++
			}
		}
		switch accumulator {
		case 3:
			totalWinnings += 1
		case 4:
			totalWinnings += 10
		case 5:
			totalWinnings += 100
		case 6:
			totalWinnings += 1000
		}
	}
	return totalWinnings
}

func searchTicket(num int, winNum []int) bool {
	start, end := 0, len(winNum)-1
	for end >= 0 {
		middle := (start + end) / 2
		if winNum[middle] == num {
			return true
		}
		if start == end {
			return false
		}
		if winNum[middle] < num {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
