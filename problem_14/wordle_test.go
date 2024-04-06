package wordle

import "testing"

var TestCases = []struct {
	Description string
	Guesses     [][]string
	Expected    string
}{
	{
		"Base Test Case",
		[][]string{
			{"hapless", "GBYYYBB"},
			{"jackpot", "BBBBYBB"},
			{"fullest", "YYGYYBB"},
		},
		"helpful",
	},
	{
		"Real Test Case",
		[][]string{
			{"keyless", "YYBBYYG"},
			{"society", "YGYYYBB"},
			{"phobias", "BBGBGBG"},
		},
		"cookies",
	},
}

func TestGuessWord(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			pof := NewProfile(test.Guesses)
			word, err := pof.GuessWord("word_list.txt")
			if err != nil {
				t.Fatal(err)
			}
			if word != test.Expected {
				t.Fatalf("expected: %s, got: %s", test.Expected, word)
			}

		})
	}
}
