package main

import (
	"bufio"
	"fmt"
	"os"
)

type All struct {
}

func main() {

	file, _ := os.Open("input_sample")
	// file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsedInput := parseInput(scanner)

	task1Result := task1(parsedInput)
	task2Result := task2(parsedInput)

	fmt.Printf("Task 1 result: %v\n", task1Result)
	fmt.Printf("Task 2 result: %v\n", task2Result)

}

func parseInput(scanner *bufio.Scanner) All {
	input := All{}
	for scanner.Scan() {
		line := scanner.Text()

	}

	return input
}

func task1(parsedInput All) int {
	sum := 0

	return sum
}

func task2(parsedInput All) int {
	sum := 0

	return sum
}
