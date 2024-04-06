package transmission

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    int64
}{
	{
		"Base Test Case",
		"base_case.txt",
		3072,
	},
	{
		"Real Test Case",
		"real_case.txt",
		297,
	},
}

func createData(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	data := make([][]string, 0)
	for scanner.Scan() {
		data = append(data, strings.Fields(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func TestValidateAndCorrect(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			data, err := createData(test.InputFile)
			if err != nil {
				t.Fatal("data generation failed, faulty tests")
			}
			got, err := ValidateAndCorrect(data)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}

func BenchmarkValidateAndCorrect(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}
	data, err := createData(TestCases[1].InputFile)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		ValidateAndCorrect(data)
	}
}
