package message

import "testing"

var TestCases = []struct {
	Description string
	DecodeFile  string
	Input       string
	Expected    string
}{
	{
		"Base Test Case",
		"decode_file.txt",
		"c6ab1f512c3fff",
		"GOOD_DAY",
	},
	{
		"Real Test Case",
		"decode_file.txt",
		"1724cf8567eb02c3d384b21a63f588c5fd0b87a65c03ea4534a3fbf47e9d7207ac2b3f409a570847d18fbf585670f58002394a99fd0afa3ae482e60f42e1eb24172c82ade8923eb2488702c2beb11ca5eb287aace8fe8b2648050c8423efd6236734c7eb0acfd0660861f4d3a5019bd1a7ac2b387aa2d724c421e9a860940f582633807bf4afe8923e85c3d12362087a36f5d397aace427924611f5970fae9cbdfa80632845bd6429cbd6529d71940413dfff",
		"THANKYOU_FOR_HELPING_WITH_THE_SOFT_LAUNCH._I_HOPE_YOU_ENJOYED_IT._YOUR_FEEDBACK_TO_IMPROVE_THE_PROBLEMS_AND_PLATFORM_WILL_BE_MUCH_APPRECIATED._WISHING_YOU_EVERY_SUCCESS_IN_YOUR_COMPUTER_SCIENCE_FUTURE._OH_AND_THE_ANSWER_IS_42_MULTIPLIED_BY_42._REGARDS_PAUL_BAUMGARTEN.",
	},
}

func TestDecode(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := Decode(test.DecodeFile, test.Input)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("wanted: %s, got: %s", test.Expected, got)
			}
		})
	}
}

var HexToBinTestCases = []struct {
	Description string
	Input       string
	Expected    string
}{
	{
		"First test",
		"c6ab1f512c3fff",
		"11000110101010110001111101010001001011000011111111111111",
	},
}

func TestHexToBin(t *testing.T) {
	for _, test := range HexToBinTestCases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := HexToBin(test.Input)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("expected: %s, got: %s", test.Expected, got)
			}
		})
	}
}

func BenchmarkHexToBin(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}
	for i := 0; i < b.N; i++ {
		HexToBin(HexToBinTestCases[0].Input)
	}
}
