package lottery

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var TestCases = []struct {
	Description    string
	WinningNumbers []int
	TicketsFile    string
	Expected       int
}{
	{
		"Base Test Case",
		[]int{12, 48, 30, 95, 15, 55, 97},
		"base_case.txt",
		110,
	},
	{
		"Real Test Case",
		[]int{12, 48, 30, 95, 15, 55, 97},
		"real_case.txt",
		56,
	},
}

func parseTickets(filePath string) ([][]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	tickets := make([][]int, 0)
	for scanner.Scan() {
		ticket := make([]int, 0, 6)
		for _, numChar := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			ticket = append(ticket, num)
		}
		tickets = append(tickets, ticket)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tickets, nil
}

func TestCalculateWinnings(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			tickets, err := parseTickets(test.TicketsFile)
			if err != nil {
				t.Fatalf("unable to run test, got following error while paring tickets: %s", err.Error())
			}
			got := CalculateWinnings(tickets, test.WinningNumbers)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}
