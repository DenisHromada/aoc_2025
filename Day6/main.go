package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type All struct {
	lines     [][]int
	operators []byte
}

func main() {

	// file, _ := os.Open("input_sample")
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsedInput := parseInput(scanner)

	// fmt.Printf(string(parsedInput.operators))

	task1Result := task1(parsedInput)
	task2Result := task2(parsedInput)

	fmt.Printf("Task 1 result: %v\n", task1Result)
	fmt.Printf("Task 2 result: %v\n", task2Result)

}

func parseInput(scanner *bufio.Scanner) All {
	input := All{}

	for scanner.Scan() {
		line := scanner.Text()
		slices := strings.Fields(line)
		lineNumbers := make([]int, 0)

		isOps := false

		for _, el := range slices {
			num, err := strconv.Atoi(el)
			if err != nil {
				input.operators = append(input.operators, el[0])
				isOps = true
			} else {
				lineNumbers = append(lineNumbers, num)
			}
		}
		if !isOps {
			input.lines = append(input.lines, lineNumbers)
		} else {
			break
		}

	}

	for _, line := range input.lines {
		for _, el := range line {
			fmt.Printf("%v ", el)
		}
		fmt.Printf("\n")
	}
	for _, el := range input.operators {
		fmt.Printf("%v ", el)
	}
	fmt.Printf("\n")

	return input
}

func task1(parsedInput All) int {
	sum := 0

	results := make([]int, len(parsedInput.operators))
	for i, op := range parsedInput.operators {
		for j, num := range parsedInput.lines {
			if j == 0 {
				results[i] = parsedInput.lines[j][i]
				continue
			}
			switch op {
			case '*':
				results[i] *= num[i]
			case '+':
				results[i] += num[i]
			}
		}
	}

	fmt.Printf("%v\n", results)
	for _, num := range results {
		sum += num
	}

	return sum
}

func task2(parsedInput All) int {
	sum := 0

	return sum
}
