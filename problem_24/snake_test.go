package snake

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description string
	FruitFile   string
	Moves       string
	Expected    int
}{
	{
		"Base case Test",
		"base_test.txt",
		"DDDRRRDDLLLDRRRRRRRDD",
		320,
	},
}

func parseFruits(inputFile string) ([]Position, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]Position, 0)
	for scanner.Scan() {
		splitText := strings.Fields(scanner.Text())
		for _, text := range splitText {
			fruitCord := strings.Split(text, ",")
			x, err := strconv.Atoi(fruitCord[0])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(fruitCord[1])
			if err != nil {
				return nil, err
			}
			fruit := Position{
				X: x,
				Y: y,
			}
			out = append(out, fruit)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestCalculateScore(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			fruits, err := parseFruits(tt.FruitFile)
			if err != nil {
				t.Error(err)
			}
			got := CalculateScore(fruits, tt.Moves)
			if tt.Expected != got {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}
