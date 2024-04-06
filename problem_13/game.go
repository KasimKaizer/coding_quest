package snakeladder

// define turn type containing moves from player 1 and two, in a single turn
// create a gameboard type representing the grid
// create a function to parse the grid, probably done in the test file
// create a method on the gameboard called move, this method would move the piece to
// its new position
// remember even: straight, odd: backwards.

type GameBoard [][]int

type Turn struct {
	player1 int
	player2 int
}

type position struct {
	col int
	row int
}

func (g GameBoard) Play(turns []Turn) int {
	positions := []position{
		{0, 0},
		{0, 0},
	}

	for turnNum, turn := range turns {
		for player, roll := range []int{turn.player1, turn.player2} {
			pos := positions[player]
			for roll != 0 {
				if roll > 0 {
					pos = g.moveForward(pos, roll)
				} else {
					pos = g.moveBackwards(pos, roll*-1)
				}
				if pos.row > len(g)-1 || pos.row == len(g)-1 && pos.col == 0 {
					return (player + 1) * (turnNum + 1)
				}
				roll = g[pos.row][pos.col]
			}
			positions[player] = pos

		}
	}
	return -1
}

func (g GameBoard) moveForward(pos position, move int) position {
	// create a type to demo the position of a player on the gameboard
	// lets say the position we are at is (1,3), so to move 10 steps on a 10x10 board
	// 3+10%10: new col pos, row + 3+10/10: new row pos
	// if odd then (9-3)+10%10: new col pos, (9-3)+10/10: new row pos
	// then the new row is odd then just convert it back to opposite pos by
	// subtracting 9 from the col pos
	// handle negative input
	col := (pos.col + move) % len(g[pos.row])
	step := (pos.col + move) / len(g[pos.row])

	if pos.row%2 != 0 {
		col = ((len(g[pos.row]) - 1) - pos.col + move) % len(g[pos.row])
		step = (((len(g[pos.row]) - 1) - pos.col) + move) / len(g[pos.row])

	}
	row := pos.row + step
	if row%2 != 0 {
		col = (len(g[row]) - 1) - col
	}
	return position{col, row}
}

func (g GameBoard) moveBackwards(pos position, move int) position {
	col := (((len(g[pos.row]) - 1) - pos.col) + move) % len(g[pos.row])
	step := (((len(g[pos.row]) - 1) - pos.col) + move) / len(g[pos.row])

	if pos.row%2 != 0 {
		col = (pos.col + move) % len(g[pos.row])
		step = (pos.col + move) / len(g[pos.row])

	}
	row := pos.row - step
	if row%2 == 0 {
		col = (len(g[row]) - 1) - col
	}
	return position{col, row}
}
