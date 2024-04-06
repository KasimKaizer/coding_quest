package forgery

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// sha256.New().write([]byte(line)) , this will return a byte slice,
// a func to verify records, it returns a bool to indicate if records are tempered with and a int
// which represents the index where the record was tampered.

func Validate(dataSet []string) (bool, int, error) {
	// instead of returning bool, why not just return -1 as index if everything is ok
	for idx, data := range dataSet {
		sep := strings.LastIndex(data, "|")
		h := sha256.New()
		_, err := h.Write([]byte(data[:sep]))
		if err != nil {
			return false, 0, err
		}
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if hash != data[sep+1:] {
			return false, idx, nil
		}
	}
	return true, 0, nil
}

func Correct(corruptedData []string) ([]string, error) {
	//just accept the whole dataset with the idx of the bad item, we can work from there.
	// this is also not a ideal solution
	// TODO: refactor this function
	ok, _, err := Validate(corruptedData[:1])
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("forgery.Correct: the first item in the provided data need to be correct")
	}
	prevHash := corruptedData[0][strings.LastIndex(corruptedData[0], "|")+1:]
	out := []string{corruptedData[0]}
	for _, data := range corruptedData[1:] {
		splitData := strings.Split(data, "|")
		minedNum, err := strconv.Atoi(splitData[1])
		if err != nil {
			return nil, err
		}
		hash, initial := "", true
		for !strings.HasPrefix(hash, "000000") {
			h := sha256.New()
			_, err := h.Write([]byte(fmt.Sprintf("%s|%d|%s", splitData[0], minedNum, prevHash)))
			if err != nil {
				return nil, err
			}
			hash = fmt.Sprintf("%x", h.Sum(nil))
			minedNum++
			if initial {
				// if the initial was correct then this would just make our
				// minedNum 0, which will then become -1, fix this or maybe just remove this as a whole
				minedNum = 0
				initial = false
			}
		}
		fixedData := fmt.Sprintf("%s|%d|%s|%s", splitData[0], minedNum-1, prevHash, hash)
		prevHash = hash
		out = append(out, fixedData)
	}
	return out, nil
}

func ValidateAndCorrect(data []string) ([]string, error) {
	ok, idx, err := Validate(data)
	if err != nil {
		return nil, err
	}
	if ok {
		return data, nil
	}
	proper, err := Correct(data[idx-1:])
	// this is a very bad implementation.
	// better implementation would be to just modify the og slice, which was passed.
	// as we don't really care about the data being modified. i have not because of the tests
	if err != nil {
		return nil, err
	}
	proper = append(data[:idx-1], proper...)
	return proper, nil
}
