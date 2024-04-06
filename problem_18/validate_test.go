package validate

import "testing"

var TestCase = []struct {
	Description string
	InputFile   string
	Expected    uint
}{
	{
		"Base Case Test",
		"base_test.txt",
		187733700,
	},
	{
		"Real Case Test",
		"real_test.txt",
		13327755200,
	},
}

func TestValidateInventory(t *testing.T) {
	for _, tt := range TestCase {
		t.Run(tt.Description, func(t *testing.T) {
			got, err := ValidateInventory(tt.InputFile)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Expected {
				t.Fatalf("expected: %d, got: %d", tt.Expected, got)
			}
		})
	}
}

func BenchmarkValidateInventory(b *testing.B) {
	for _, tt := range TestCase {
		b.Run(tt.Description, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ValidateInventory(tt.InputFile)
			}
		})
	}
}
