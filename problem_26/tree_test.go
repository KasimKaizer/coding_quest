package tree

import (
	"bufio"
	"os"
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
		16,
	},
	{
		"Real Case Test",
		"real_test.txt",
		30784,
	},
}

func parseData(inputPath string) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestMaxWidthAndHeight(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			tree, err := NewTree(data)
			if err != nil {
				t.Error(err)
			}
			width, height := tree.MaxWidthAndHeight()
			if tt.Expected != (width * height) {
				t.Fatalf("expected: %d, got: %d", tt.Expected, (width * height))
			}
		})
	}
}
