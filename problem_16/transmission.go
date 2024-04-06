package transmission

import (
	"errors"
	"strconv"
)

// convert hex to decimal, go through rows, calculating until we hit a row where sum doesn't match
// when we hit it, start calculating columns, again until we hit a col where sum doesn't match
// we found the wrong hex, calculate difference of proper row and col from what we got,
// minus it from the wrong hex, and viola (probably wrong spelling) we have the correct hex

// we would let the solution parse hex, the test would just convert hex into a matrix
// sum % 256
type recursionMode int

const (
	rows recursionMode = iota
	columns
)

func ValidateAndCorrect(data [][]string) (int64, error) {
	lastRow := len(data) - 1
	lastCol := len(data[0]) - 1
	errRow, errRowDiff, err := calculateSum(data, lastRow, lastCol, rows)
	if err != nil {
		return 0, err
	}
	errCol, errColDiff, err := calculateSum(data, lastCol, lastRow, columns)
	if err != nil {
		return 0, err
	}
	if errRowDiff != errColDiff {
		return 0, errors.New("transmission.ValidateAndCorrect: multiple erroneous bytes in data")
	}
	errByte, err := strconv.ParseInt(data[errRow][errCol], 16, 64)
	if err != nil {
		return 0, err
	}
	return errByte * (errByte - errRowDiff), nil
}

func calculateSum(data [][]string, outer, inner int, mode recursionMode) (int, int64, error) {
	for i := 0; i < outer; i++ {
		sum := int64(0)
		for j := 0; j < inner; j++ {
			var num int64
			var err error
			switch mode {
			case 0:
				num, err = strconv.ParseInt(data[i][j], 16, 64)
			case 1:
				num, err = strconv.ParseInt(data[j][i], 16, 64)
			}
			if err != nil {
				return 0, 0, err
			}
			sum += num
		}
		var check int64
		var err error
		switch mode {
		case 0:
			check, err = strconv.ParseInt(data[i][inner], 16, 64)
		case 1:
			check, err = strconv.ParseInt(data[inner][i], 16, 64)
		}
		if err != nil {
			return 0, 0, err
		}
		sum %= 256
		if sum == check {
			continue
		}
		diff := sum - check
		if diff < 0 {
			diff += 256
		}
		return i, diff, nil
	}
	return -1, -1, nil
}
