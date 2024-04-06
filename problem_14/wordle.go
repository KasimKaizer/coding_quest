package wordle

import (
	"bufio"
	"os"
)

// create a function to create "a profile with guesses", use the guess and its hint to construct
// a profile which we can use to then figure out the word, we go through the whole 7 word list, matching the pattern
// should not be that hard, more of a tricky exercise, creating a profile might be kinda difficult but its fine

type Profile struct {
	NotContain map[byte]bool
	Contain    map[byte]bool
	Perfect    []byte
}

func NewProfile(guesses [][]string) *Profile {
	newProf := new(Profile)
	newProf.Contain = make(map[byte]bool, 7)
	newProf.NotContain = make(map[byte]bool, 7)
	newProf.Perfect = make([]byte, 7)
	for _, guess := range guesses {
		for idx, char := range guess[1] {
			switch char {
			case 'B':
				newProf.NotContain[guess[0][idx]] = true
			case 'G':
				newProf.Perfect[idx] = guess[0][idx]
				fallthrough
			case 'Y':
				newProf.Contain[guess[0][idx]] = true

			}
		}

	}
	return newProf
}

func (p *Profile) GuessWord(wordList string) (string, error) {
	f, err := os.Open(wordList)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ok := p.CheckWord(scanner.Text())
		if ok {
			return scanner.Text(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "notfound", nil
}

func (p *Profile) CheckWord(word string) bool {
	for idx, char := range []byte(word) {
		if p.Perfect[idx] != '\x00' {
			if p.Perfect[idx] != char {
				break
			}
		}
		if _, ok := p.NotContain[char]; ok {
			break
		}
		if _, ok := p.Contain[char]; !ok {
			break
		}
		if idx == 6 {
			return true
		}
	}
	return false
}
