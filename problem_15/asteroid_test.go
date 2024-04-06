package asteroid

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
		8,
	},
	{
		"Real Test Case",
		"real_test.txt",
		33,
	},
}

func TestCalAverMass(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			grid, err := constructGrid(test.InputFile)
			if err != nil {
				t.Fatal("failed while trying to generate the grid, faulty tests")
			}
			got := CalAverMass(grid)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}

func constructGrid(filePath string) ([][]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([][]int, 0)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		curRow := make([]int, len(line))
		for idx, numChar := range line {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			curRow[idx] = num
		}
		out = append(out, curRow)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func BenchmarkCalAverMass(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}
	grid, err := constructGrid(TestCases[1].InputFile)
	if err != nil {
		b.Fatal("failed while trying to generate the grid, faulty tests")
	}
	for i := 0; i < b.N; i++ {
		CalAverMass(grid)
	}
}
