package vigenere

import (
	"strings"
)

type mode int

const (
	decode mode = iota + 1
	encode
)

func Encode(message, key, charSet string) string {
	return cipher(encode, message, key, charSet)
}

func Decode(message, key, charSet string) string {
	return cipher(decode, message, key, charSet)
}
func cipher(cipherMode mode, message, key, charSet string) string {
	if len(message) == 0 || len(key) == 0 || len(charSet) == 0 {
		return ""
	}
	lookup := reverseLookup(charSet)
	keyIdx := -1
	var output strings.Builder
	for _, char := range []byte(message) {
		keyIdx = (keyIdx + 1) % len(key)
		if lookup[char] == 0 && char != charSet[0] {
			output.WriteByte(char)
			continue
		}
		pos := lookup[char] + (lookup[key[keyIdx]] + 1)
		if cipherMode == decode {
			pos = lookup[char] - (lookup[key[keyIdx]] + 1)
		}
		pos = (pos%len(charSet) + len(charSet)) % len(charSet)
		output.WriteByte(charSet[pos])

	}
	return output.String()
}

func reverseLookup(charSet string) []int {
	reverseLookup := make([]int, 256)
	for idx, char := range []byte(charSet) {
		reverseLookup[char] = idx
	}
	return reverseLookup
}
