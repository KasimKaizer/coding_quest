package debugging

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ADD target source - Take the value of source and add it to target
// MOD target source - Calculate the modulus target % source, and save to target
// DIV target source - Perform integer division of target // source, and save to target
// MOV target source - Take the value of source and copy it into target (replacing whatever was there)
// JMP source - Jump source number of instructions within the code
// JIF source - Jump source number of instructions within the code IF the most recent CEQ or CGE operation was TRUE
// CEQ source1 source2 - Compare the values in source1 and source2, are they equal?
// CGE source1 source2 - Compare the values in source1 and source2. Is source1 greater than or equal to source2?
// OUT source - Output the value of source to the terminal
// END - Terminate the program

func InterpretAsm(instructions []string) error {
	wasEqual := false
	variables := make(map[string]int)
	for i := 0; i < len(instructions); i++ {
		values := strings.Fields(instructions[i])
		if len(values) == 0 {
			return errors.New("debugging.InterpretAsm: invalid code")
		}
		switch values[0] {
		case "ADD", "MOD", "DIV", "MOV":
			num, err := strconv.Atoi(values[2])
			if err != nil {
				num = variables[values[2]]
			}
			performOperation(variables, values[0], values[1], num)
		case "JMP", "JIF":
			if values[0] == "JIF" && !wasEqual {
				break
			}
			wasEqual = false
			num, err := strconv.Atoi(values[1])
			if err != nil {
				num = variables[values[1]]
			}
			i += (num - 1)
		case "CEQ", "CGE":
			value1, err := strconv.Atoi(values[1])
			if err != nil {
				value1 = variables[values[1]]
			}
			value2, err := strconv.Atoi(values[2])
			if err != nil {
				value2 = variables[values[2]]
			}
			if values[0] == "CEQ" {
				wasEqual = value1 == value2
				break
			}
			wasEqual = value1 >= value2
		case "OUT":
			num, err := strconv.Atoi(values[1])
			if err != nil {
				num = variables[values[1]]
			}
			fmt.Println(num)
		case "END":
			return nil
		}
	}
	return errors.New("debugging.InterpretAsm: program did not end")
}

func performOperation(variables map[string]int, ops, key string, value int) {
	switch ops {
	case "ADD":
		variables[key] += value
	case "MOD":
		variables[key] %= value
	case "DIV":
		variables[key] /= value
	case "MOV":
		variables[key] = value
	}
}