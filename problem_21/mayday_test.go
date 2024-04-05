package mayday

import (
	"testing"
)

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    string
}{
	{
		"Base test case",
		"base_test.txt",
		"This is a test. This is a test. Thankyou.",
	},
	{
		"Real test case",
		"real_test.txt",
		"This is the SY Titanic II. We have struck an asteroid. Engines non responsive. Please help. Coordinates 242.03 by 182.24 by 92.58. Message ends.",
	},
}

func TestDecodeMessage(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.Description, func(t *testing.T) {
			got, err := DecodeMessage(tt.InputFile)
			if err != nil {
				t.Error(err)
			}
			if got != tt.Expected {
				t.Fatalf("expected: %s, got: %s", tt.Expected, got)
			}
		})
	}
}
