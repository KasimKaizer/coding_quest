package shields

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description   string
	GridRows      int
	GridCols      int
	InputDataFile string
	Expected      int
}{
	{
		"Base Test Case",
		10,
		10,
		"base_test.txt",
		12,
	},
	{
		"Example Test Case",
		100,
		100,
		"example_test.txt",
		2061,
	},
	{
		"Real Test Case",
		100000,
		20000,
		"real_test.txt",
		154807700,
	},
}

func parseData(file string) ([][]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var out [][]int
	for scanner.Scan() {
		var temp []int
		for _, numChar := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			temp = append(temp, num)
		}
		out = append(out, temp)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func TestFindVulnerabilities(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputDataFile)
			if err != nil {
				t.Fatal(err)
			}
			got := FindVulnerabilities(data, tt.GridRows, tt.GridCols)
			if got != tt.Expected {
				t.Fatalf("got: %d, expected: %d", got, tt.Expected)
			}
		})
	}
}

func BenchmarkFindVulnerabilities(b *testing.B) {
	for _, tt := range TestCases {
		b.Run(tt.Description, func(b *testing.B) {
			data, err := parseData(tt.InputDataFile)
			if err != nil {
				b.Fatal(err)
			}
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				FindVulnerabilities(data, tt.GridRows, tt.GridCols)
			}
		})
	}
}
