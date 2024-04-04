package tictactoe

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    int
}{
	{
		"Base Test Case",
		"base_test.txt",
		0,
	},
	{
		"Real Test Case",
		"real_test.txt",
		20938290,
	},
}

func parseData(filePath string) ([][]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var games [][]int
	for scanner.Scan() {
		splitText := strings.Fields(scanner.Text())
		var game []int
		for _, numChar := range splitText {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			game = append(game, int(num))
		}
		games = append(games, game)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func TestPlayGames(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			got := PlayGames(data)
			if got != tt.Expected {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}
