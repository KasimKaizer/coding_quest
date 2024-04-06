package diagnostics

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    int
}{
	// {
	// 	"Base Test Case",
	// 	"base_data.txt",
	// 	10,
	// },
	{
		"Real Test Case",
		"real_data.txt",
		10,
	},
}

func parseData(filePath string) ([]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		output = append(output, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func TestFindInconsistences(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			data, err := parseData(test.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			got, err := FindInconsistences(data)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}
