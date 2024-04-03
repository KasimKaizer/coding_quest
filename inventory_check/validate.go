package validate

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ValidateInventory(filePath string) (uint, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	category := make(map[string]uint)
	for scanner.Scan() {
		splitText := strings.Fields(scanner.Text())
		if len(splitText) != 3 {
			return 0, errors.New("validate.ValidateInventory: invalid data provided")
		}
		num, err := strconv.Atoi(splitText[1])
		if err != nil {
			return 0, err
		}
		category[splitText[2]] += uint(num)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	product := uint(1)
	for _, value := range category {
		product *= (value % 100)
	}
	return product, nil
}
