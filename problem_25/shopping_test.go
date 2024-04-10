package shopping

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
		"Base Case Test",
		"base_test.txt",
		43,
	},
}

func generateGraph(inputFile string) ([][]int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var out [][]int
	for scanner.Scan() {
		splitText := strings.Fields(scanner.Text())
		var row []int
		for _, numChar := range splitText {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		out = append(out, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestFindShortestPath(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			graph, err := generateGraph(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			got := FindShortestPath(graph)
			if tt.Expected != got {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}
