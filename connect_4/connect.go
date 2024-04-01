package connect4

type Player int

const (
	Player1 Player = iota + 1
	Player2
	Player3
)

// to get next player curPlayer = curPlayer + 1 % len(players)
// do depth first search after each players plays their turn to determine a winner
// moves are as string of long numbers, so (moves[idx] - '0' -1 ) to get the column
// it would be better to have
// win determination in a separate function
// board would be of []Player with len of 7, we create a board for each game.
// i have decided to have separate function for the a game played after all.
// much cleaner code that way
// travel in 8 directions, can make it tricky but not hard.
// a massive oversight on my part, the win doesn't need to be from the current position.
// it can be from any position

var directions = [...][2]int{
	{-1, 0},  // top
	{-1, 1},  // top-right
	{0, 1},   // right
	{1, 1},   // bottom-right
	{1, 0},   // bottom
	{1, -1},  // bottom-left
	{0, -1},  // left
	{-1, -1}, // top-left
}

func DetermineResult(games []string) int {
	play1Wins, play2Wins, play3Wins := 0, 0, 0
	for _, game := range games {
		winner := DetermineGame(game)
		switch winner {
		case Player1:
			play1Wins++
		case Player2:
			play2Wins++
		case Player3:
			play3Wins++
		}
	}
	return (play1Wins * play2Wins * play3Wins)
}

func DetermineGame(moves string) Player {
	board := make([][]Player, 7)
	players := [...]Player{Player1, Player2, Player3}
	curPlayer := 0
	for _, move := range []byte(moves) {
		row := int((move - '0') - 1)
		board[row] = append(board[row], players[curPlayer])
		if IsWin(board, players[curPlayer], row, (len(board[row]) - 1)) {
			return players[curPlayer]
		}
		curPlayer = (curPlayer + 1) % len(players)
	}
	return 0
}

func IsWin(board [][]Player, player Player, row, col int) bool {
	toVisit := [][]int{{row, col}}
	first := true
	for i := 0; i < len(toVisit); i++ {
		for _, direct := range directions {
			newR, newC := toVisit[i][0], toVisit[i][1]
			for i := 1; i <= 3; i++ {
				newR, newC = newR+direct[0], newC+direct[1]
				if newR >= len(board) || newR < 0 {
					break
				}
				if newC >= len(board[newR]) || newC < 0 {
					break
				}
				if board[newR][newC] != player {
					break
				}
				if i == 3 {
					return true
				}
				if first {
					toVisit = append(toVisit, []int{newR, newC})
				}
			}
		}
		first = false
	}

	return false
}
