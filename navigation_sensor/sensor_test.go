package sensor

import (
	"fmt"
	"testing"
)

var ParseBitFlipTestCases = []struct {
	Description string
	InputFile   string
	Expected    int
}{
	{
		"Base Test Case",
		"base_test.txt",
		17837,
	},
	{
		"Real Test Case",
		"real_case.txt",
		297,
	},
}

var IsEvenParityTestCases = []struct {
	Input    int
	Expected bool
}{
	{30635, false},
	{34132, true},
	{46818, false},
	{31114, true},
	{53800, true},
}

func TestParseBitFlip(t *testing.T) {
	for _, tt := range ParseBitFlipTestCases {
		t.Run(tt.Description, func(t *testing.T) {
			got, err := ParseBitFlip(tt.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Expected {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}

func TestIsEvenParity(t *testing.T) {
	for _, tt := range IsEvenParityTestCases {
		t.Run(fmt.Sprintf("Test for: %d", tt.Input), func(t *testing.T) {
			got := IsEvenParity(uint16(tt.Input))
			if got != tt.Expected {
				t.Fatalf("expected: %v, got: %v", tt.Expected, got)
			}
		})
	}
}
