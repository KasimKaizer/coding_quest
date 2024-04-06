package debugging

import (
	"bufio"
	"log"
	"os"
)

func parseFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	out := make([]string, 0)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func ExampleInterpretAsm() {
	data, err := parseFile("base_code.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = InterpretAsm(data)
	if err != nil {
		log.Fatal(err)
	}
	// Output: 5
}

func ExampleInterpretAsm_second() {
	data, err := parseFile("example_code.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = InterpretAsm(data)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
}

func ExampleInterpretAsm_third() {
	data, err := parseFile("real_code.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = InterpretAsm(data)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// 7745743850156157
}
