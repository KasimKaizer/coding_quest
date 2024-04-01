package forgery

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

var ValidateTestCase = []struct {
	Description   string
	Input         []string
	Expected      bool
	ExpectedIndex int
}{
	{
		"Base Test Case",
		[]string{"Original iPhone still in box|3595421|0000000000000000000000000000000000000000000000000000000000000000|00000078f97879b26be6baf2adb520b126f84ed10464ed4e9a77b8ed87e07468"},
		true,
		0,
	},
}

var CorrectTestCases = []struct {
	Description string
	Input       []string
	Expected    []string
}{
	{
		"Base Test Case",
		[]string{
			"Original iPhone still in box|3595421|0000000000000000000000000000000000000000000000000000000000000000|00000078f97879b26be6baf2adb520b126f84ed10464ed4e9a77b8ed87e07468",
			"Apollo 11 moon rock|27703084|00000078f97879b26be6baf2adb520b126f84ed10464ed4e9a77b8ed87e07468|00000068a1928374565849384594855858566495b6ac56266edb0389c2d9a045",
		},
		[]string{
			"Original iPhone still in box|3595421|0000000000000000000000000000000000000000000000000000000000000000|00000078f97879b26be6baf2adb520b126f84ed10464ed4e9a77b8ed87e07468",
			"Apollo 11 moon rock|27703084|00000078f97879b26be6baf2adb520b126f84ed10464ed4e9a77b8ed87e07468|00000068a1e823c97e72ff22b0450dc4cfa66495b6ac56266edb0389c2d9a045",
		},
	},
}

func parseData(testDataPath string) ([]string, error) {
	f, err := os.Open(testDataPath)
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
	return out, err
}

func TestValidate(t *testing.T) {
	for _, test := range ValidateTestCase {
		t.Run(test.Description, func(t *testing.T) {
			got, idx, err := Validate(test.Input)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("expected: %t, got: %t", test.Expected, got)
			}
			if idx != test.ExpectedIndex {
				t.Fatalf("expected index: %d, got index: %d", test.ExpectedIndex, idx)
			}
		})
	}
}

func TestCorrect(t *testing.T) {
	for _, test := range CorrectTestCases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := Correct(test.Input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(test.Expected, got) {
				t.Fatalf("expected: %v, got: %v", test.Expected, got)
			}
		})
	}
}
