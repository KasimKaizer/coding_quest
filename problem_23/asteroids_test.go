package asteroids

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
	GridWidth   int
	GridHeight  int
	Expected    string
}{
	{
		"Base Test Cases",
		"base_test.txt",
		8,
		8,
		"5:4",
	},
	{
		"Real Test Cases",
		"real_test.txt",
		100,
		100,
		"5:8",
	},
}

func parseData(filePath string) ([]Asteroid, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := make([]Asteroid, 0)
	for scanner.Scan() {
		splitNum := strings.Fields(scanner.Text())
		var temp []float64
		for _, charNum := range splitNum {
			num, err := strconv.ParseFloat(charNum, 64)
			if err != nil {
				return nil, err
			}
			temp = append(temp, num)
		}
		asteroid := Asteroid{
			XPos:   temp[0],
			YPos:   temp[1],
			XSpeed: temp[2],
			YSpeed: temp[3],
		}
		output = append(output, asteroid)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func TestFindSafePath(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			got := FindSafePath(data, tt.GridWidth, tt.GridHeight)
			if tt.Expected != got {
				t.Fatalf("expected: %s,got: %s", tt.Expected, got)
			}
		})
	}
}
