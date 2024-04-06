package vigenere

import "testing"

var DecodeTestCases = []struct {
	Description string
	Input       string
	InputChars  string
	InputKey    string
	Expected    string
}{
	{
		"Base Case Test",
		"PJ UTGX LF JXFFW",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"SECRET",
		"WE COME IN PEACE",
	},
	{
		"First Example Case Test",
		"dAyevvMbfHgENFsy:fDqnGddIzfMqm",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"With great power comes great responsibility",
		"Are you enjoying coding quest?",
	},
	{
		"Second Example Case Test",
		"lfwwrsvbvMbmIEnK:wDjutpzoxfwowypDDHxB(rzfwKXBMd",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"With great power comes great responsibility",
		"I could use this to pass secret notes in class!",
	},
	{
		"Real Case Test",
		"ftmpH.:lemGubTDmMb'YtfsublbnkKlMmOoKywmmOIpa.,3mNeEbl?(bVtkUy?xtoNtCkAg:;n)OlInqp2rjap6JwiG)9H'jHm: pjok'9njQbtOxusdql'b'VtkrBb5j!aMWGieIjOHfrw,j,ubsbm,xrufoKljGdob8q,APzqI:0fpi:.Jsipk6lueD):!wrwbd?j(LbmODCCz7:vjbANCsqp2ts);Of,?p; lulx,tXGbLmbTflKBbYlCCdle1bnYtGrCl1bnw:PrphBeYFviLoZD.7pb!)nrztr0lCvl8n'tqIHn8",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"Roads? Where We're Going, We Don't Need Roads.",
		"Hello! I hope you are enjoying your trip to Ral'Malgor. Don't forget to pick up some souvenirs for the family while you are there. Perhaps send mom a postcard as well? Also make sure to take some great photos! See you soon!! ... by the way the answer to the question is 'codingquest2022' (without the quotes).",
	},
}

var EncodeTestCases = []struct {
	Description string
	Input       string
	InputChars  string
	InputKey    string
	Expected    string
}{
	{
		"Base Case Test",
		"WE COME IN PEACE",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"SECRET",
		"PJ UTGX LF JXFFW",
	},
	{
		"First Example Case Test",
		"Are you enjoying coding quest?",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"With great power comes great responsibility",
		"dAyevvMbfHgENFsy:fDqnGddIzfMqm",
	},
	{
		"Second Example Case Test",
		"I could use this to pass secret notes in class!",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"With great power comes great responsibility",
		"lfwwrsvbvMbmIEnK:wDjutpzoxfwowypDDHxB(rzfwKXBMd",
	},
	{
		"Real Case Test",
		"Hello! I hope you are enjoying your trip to Ral'Malgor. Don't forget to pick up some souvenirs for the family while you are there. Perhaps send mom a postcard as well? Also make sure to take some great photos! See you soon!! ... by the way the answer to the question is 'codingquest2022' (without the quotes).",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,;:?! '()",
		"Roads? Where We're Going, We Don't Need Roads.",
		"ftmpH.:lemGubTDmMb'YtfsublbnkKlMmOoKywmmOIpa.,3mNeEbl?(bVtkUy?xtoNtCkAg:;n)OlInqp2rjap6JwiG)9H'jHm: pjok'9njQbtOxusdql'b'VtkrBb5j!aMWGieIjOHfrw,j,ubsbm,xrufoKljGdob8q,APzqI:0fpi:.Jsipk6lueD):!wrwbd?j(LbmODCCz7:vjbANCsqp2ts);Of,?p; lulx,tXGbLmbTflKBbYlCCdle1bnYtGrCl1bnw:PrphBeYFviLoZD.7pb!)nrztr0lCvl8n'tqIHn8",
	},
}

func TestDecode(t *testing.T) {
	for _, tt := range DecodeTestCases {
		t.Run(tt.Description, func(t *testing.T) {
			got := Decode(tt.Input, tt.InputKey, tt.InputChars)
			if got != tt.Expected {
				t.Fatalf("got: %s, expected: %s", got, tt.Expected)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for _, tt := range EncodeTestCases {
		t.Run(tt.Description, func(t *testing.T) {
			got := Encode(tt.Input, tt.InputKey, tt.InputChars)
			if got != tt.Expected {
				t.Fatalf("got: %s, expected: %s", got, tt.Expected)
			}
		})
	}
}
