package route

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Start       string
	End         string
	StopTime    int
	Expected    int
}{
	{
		"Base Case Test",
		"base_test.txt",
		"AAA",
		"ZZZ",
		10,
		115,
	},
	{
		"Real Case Test",
		"real_test.txt",
		"TYC",
		"EAR",
		600,
		165127,
	},
}

func parseData(filePath string) (map[string]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make(map[string]string)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " => ")
		out[split[0]] = split[1]
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestFindShortestPath(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			got, err := FindShortestPath(data, tt.Start, tt.End, tt.StopTime)
			if err != nil {
				t.Error(err)
			}
			if tt.Expected != got {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}
