package decodepixel

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    string
	// TODO: find a way to test
	// maybe we can keep the solution in a .txt file, read it into a bytes buffer and
	// compare buffer? for now, I will keep it as a string
}{
	{
		"Base Test Case",
		"base_test.txt",
		` 
........
.######.
.#......
.#......
.######.
.#......
.#......
.######.
`, // remember to ignore the first char in the string, when creating the tests,
		// as first char is \n which was added for better illustration of the test
	},
}

func parseData(InputFile string) ([][]int, error) {
	f, err := os.Open(InputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var output [][]int
	for scanner.Scan() {
		splitText := strings.Fields(scanner.Text())
		var row []int
		for _, numChar := range splitText {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		output = append(output, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func TestDecode(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			data, err := parseData(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			buf := new(bytes.Buffer)
			Decode(data, buf)
			got := buf.String()
			if tt.Expected != got {
				t.Fatalf("expected: %s, got: %s", tt.Expected, got)
			}
		})
	}
}
