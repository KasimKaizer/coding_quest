package mayday

import (
	"bufio"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"strings"
)

func DecodeMessage(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	unOrdered := make(map[int]string)
	for scanner.Scan() {
		pos, ok, err := processMessage(scanner.Text())
		if err != nil {
			return "", err
		}
		if !ok {
			continue
		}
		unOrdered[pos] = scanner.Text()[16:]
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	ordered := make([]string, len(unOrdered))
	for idx, text := range unOrdered {
		if idx >= len(unOrdered) {
			return "", errors.New("mayday.DecodeMessage: non sequential data")
		}
		ordered[idx] = text
	}
	bs, err := hex.DecodeString(strings.Join(ordered, ""))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(bs)), nil
}

func processMessage(message string) (int, bool, error) {
	if len(message) < 64 {
		return 0, false, nil
	}
	if message[:4] != "5555" {
		return 0, false, nil
	}
	pos, err := strconv.Atoi(message[12:14])
	if err != nil {
		return 0, false, err
	}
	chkSum, err := strconv.ParseInt(message[14:16], 16, 32)
	if err != nil {
		return 0, false, err
	}
	var sum int64
	for i := 16; i < len(message)-1; i = i + 2 {
		num, err := strconv.ParseInt(message[i:i+2], 16, 32)
		if err != nil {
			return 0, false, err
		}
		sum += num
	}
	if chkSum != (sum % 256) {
		return 0, false, nil
	}
	return pos, true, nil
}
