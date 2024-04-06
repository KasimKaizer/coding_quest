package market

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
		"Base Test Case",
		"base_test.txt",
		71,
	},
	{
		"Real Test Case",
		"real_test.txt",
		6723,
	},
}

func parseData(inputFile string) ([]string, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]string, 0)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestTraverseMaze(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			input, err := parseData(tt.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			got := TraverseMaze(input)
			if got != tt.Expected {
				t.Fatalf("got: %d, expected: %d", got, tt.Expected)
			}
		})
	}
}

func BenchmarkTraverseMaze(b *testing.B) {
	for _, tt := range TestCases {
		b.Run(tt.Description, func(b *testing.B) {
			input, err := parseData(tt.InputFile)
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				TraverseMaze(input)
			}
		})
	}
}
