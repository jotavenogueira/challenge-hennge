package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumPow(numbers []string) int {
	// if doesn't have any numbers, return 0
	if len(numbers) == 0 {
		return 0
	}

	y, _ := strconv.Atoi(numbers[0])

	total := 0
	if y <= 0 {
		total = y * y * y * y
	}

	// sum the current total with the rest of the numbers
	return total + sumPow(numbers[1:])
}

func readCases(reader *bufio.Scanner, results []string, index int) {
	// stop if all cases have been read
	if index == len(results) {
		return
	}

	reader.Scan()
	x, _ := strconv.Atoi(strings.TrimSpace(reader.Text()))

	reader.Scan()
	numbers := strings.Fields(reader.Text())

	if len(numbers) != x {
		results[index] = "-1"
	} else {
		sum := sumPow(numbers)
		results[index] = strconv.Itoa(sum)
	}

	readCases(reader, results, index+1)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(reader.Text()))

	results := make([]string, n)

	readCases(reader, results, 0)

	// prints after all cases have been read
	fmt.Println(strings.Join(results, "\n"))
}
