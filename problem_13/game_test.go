package snakeladder

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var testCases = []struct {
	Description string
	boardFile   string
	boardSize   int
	movesFile   string
	Expected    int
}{
	{
		"Base Test case",
		"base_board.txt",
		6,
		"base_moves.txt",
		13,
	},
	{
		"Real Test case",
		"real_board.txt",
		20,
		"real_moves.txt",
		95,
	},
}

func TestPlay(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			board, err := createBoard(test.boardFile, test.boardSize)
			if err != nil {
				t.Fatal("GameBoard creation failed, test faulty")
			}
			turns, err := parseTurns(test.movesFile)
			if err != nil {
				t.Fatal("turns parsing failed, test faulty")
			}
			got := board.Play(turns)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}

func createBoard(filePath string, size int) (GameBoard, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	board := make(GameBoard, size)
	for i := size - 1; scanner.Scan(); i-- {
		curRow := make([]int, size)
		rowNum := strings.Fields(scanner.Text())
		for idx, numChar := range rowNum {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			curRow[idx] = num
		}
		board[i] = curRow
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return board, nil
}

func parseTurns(filePath string) ([]Turn, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	turns := make([]Turn, 0)
	for scanner.Scan() {
		turnChar := strings.Fields(scanner.Text())
		var curTurn Turn
		for idx, rollChar := range turnChar {
			roll, err := strconv.Atoi(rollChar)
			if err != nil {
				return nil, err
			}
			if idx == 1 {
				curTurn.player2 = roll
				continue
			}
			curTurn.player1 = roll
		}
		turns = append(turns, curTurn)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return turns, nil
}
