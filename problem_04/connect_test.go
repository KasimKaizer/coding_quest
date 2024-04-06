package connect4

import (
	"bufio"
	"os"
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
		21630678,
	},
}

func parseData(inputPath string) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]string, 0)
	for scanner.Scan() {
		out = append(out, strings.TrimSuffix(scanner.Text(), "\n"))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestDetermineResult(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			data, err := parseData(test.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			got := DetermineResult(data)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}
