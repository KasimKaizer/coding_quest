package travel

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
		85,
	},
	{
		"Real Test Case",
		"real_test.txt",
		64579603,
	},
}

func parseData(filePath string) ([]*Coordinate, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]*Coordinate, 0)
	for scanner.Scan() {
		temp := make([]int, 0)
		for _, numChar := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			temp = append(temp, num)
		}
		cord := new(Coordinate)
		cord.X, cord.Y, cord.Z = temp[0], temp[1], temp[2]
		out = append(out, cord)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestCalculateDistance(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			data, err := parseData(test.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			got := CalculateDistance(data)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}
