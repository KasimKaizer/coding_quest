package market

type position struct {
	row, col int
}

func TraverseMaze(maze []string) int {
	directions := []position{
		{-1, 0}, // top
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	startCol := findEntrance(maze[0])
	if startCol == -1 {
		return -1
	}
	toVisit := []position{{0, startCol}}

	// we start traveled as 1 to account for our current starting position.
	// curNodes are started as 1 as there is only one entrance to this maze
	traveled, curNodes, nextNodes := 1, 1, 0
	mem := make(map[position]struct{})
	for i := 0; i < len(toVisit); i++ {
		cur := toVisit[i]
		if cur.row == len(maze)-1 && maze[cur.row][cur.col] == ' ' {
			return traveled
		}
		for _, d := range directions {
			new := position{
				row: (cur.row + d.row),
				col: (cur.col + d.col),
			}
			if new.row < 0 || new.col < 0 ||
				new.row > (len(maze)-1) || new.col > (len(maze[new.row])-1) {
				continue
			}
			if _, ok := mem[new]; ok || maze[new.row][new.col] != ' ' {
				continue
			}
			toVisit = append(toVisit, new)
			mem[new] = struct{}{}
			nextNodes++
		}
		curNodes--
		if curNodes == 0 {
			traveled++
			curNodes, nextNodes = nextNodes, 0
		}
	}
	return -1
}

func findEntrance(firstRow string) int {
	for i := 0; i < len(firstRow); i++ {
		if firstRow[i] == ' ' {
			return i
		}
	}
	return -1
}
