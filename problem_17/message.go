package message

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// create a start and end index, slide through the binary
// remember smallest possible binary byte is of size 4
// so start := 0, end := 3
// start++ if not found in map and
// start has 1 means, we need to slide more
// side till map returns something, then move on to next part

func Decode(decodeFile, input string) (string, error) {
	const minLen = 4
	decoder, err := parseDecoder(decodeFile)
	if err != nil {
		return "", err
	}
	inputBin, err := HexToBin(input)
	if err != nil {
		return "", err
	}
	var output strings.Builder
	start, end := 0, minLen
	for end < len(inputBin)-3 {
		char, ok := decoder[inputBin[start:end]]
		for !ok {
			end++
			char, ok = decoder[inputBin[start:end]]
		}
		if char == "*" {
			break
		}
		output.WriteString(char)
		start, end = end, end+minLen
	}

	return output.String(), nil
}

func parseDecoder(decodeFile string) (map[string]string, error) {
	f, err := os.Open(decodeFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := make(map[string]string)
	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		if len(data) > 2 {
			return nil, errors.New("message.parseDecoder: the decode file contains improper number of values")
		}
		output[data[1]] = data[0]
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func HexToBin(hex string) (string, error) {
	table := [256]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'a': "1010",
		'b': "1011",
		'c': "1100",
		'd': "1101",
		'e': "1110",
		'f': "1111",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	var output strings.Builder
	for i := range len(hex) {
		b := table[hex[i]]
		if b == "" {
			return "", errors.New("message.hexToBin: unable to parse hex")
		}
		output.WriteString(b)
	}
	return output.String(), nil
}
