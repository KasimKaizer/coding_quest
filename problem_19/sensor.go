package sensor

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func ParseBitFlip(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum, count := 0, 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		if !IsEvenParity(uint16(num)) {
			continue
		}
		sum += (num &^ (1 << 15))
		count++
	}
	return intDivideRound(sum, count), nil
}

func intDivideRound(dividend, divisor int) int {
	return int(math.Round(float64(dividend) / float64(divisor)))
}

func IsEvenParity(num uint16) bool {
	// this would only work with 16 bit unsigned integers.
	num ^= (num >> 1)
	num ^= (num >> 2)
	num ^= (num >> 4)
	num ^= (num >> 8)

	if num&1 == 1 {
		return false
	}
	return true
}
