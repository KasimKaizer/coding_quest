package tictactoe

type player int

const (
	xPlayer = iota + 1
	OPlayer
)

var directions = [][]int{
	// {row, col}
	{-1, 0}, // up
	{1, 0},  //down

	{0, 1},  // right
	{0, -1}, // left

	{-1, -1}, // top left
	{1, 1},   //bottom right

	{-1, 1}, // top right
	{1, -1}, // bottom left
}

func PlayGames(games [][]int) int {
	xWins, oWins, draws := 0, 0, 0
	for _, game := range games {
		winner := playGame(game)
		switch winner {
		case xPlayer:
			xWins++
		case OPlayer:
			oWins++
		default:
			draws++
		}
	}
	return (xWins * oWins * draws)
}

func playGame(game []int) player {
	board := createBoard(3)
	players := []player{xPlayer, OPlayer}
	curPlayer := 0
	for _, move := range game {
		row, col := (move-1)/len(board), (move-1)%len(board)
		player := players[curPlayer%2]
		board[row][col] = player
		i := 0
		winCount := 1
		for _, direction := range directions {
			if i == 0 {
				winCount = 1
			}
			nRow, nCol := row, col
			for {
				nRow, nCol = nRow+direction[0], nCol+direction[1]
				if nRow < 0 || nRow >= len(board) ||
					nCol < 0 || nCol >= len(board[nRow]) {
					break
				}
				if board[nRow][nCol] != player {
					break
				}
				winCount++
				if winCount == 3 {
					return player
				}
			}
			i = (i + 1) % 2
		}
		curPlayer++
	}
	return -1
}

func createBoard(size int) [][]player {
	out := make([][]player, size)
	for idx := range out {
		out[idx] = make([]player, size)
	}
	return out
}
