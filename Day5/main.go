package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	_ := parseInput(scanner)

	task1_result := task1(parsed_input)
	task2_result := task2(parsed_input)

	fmt.Printf("Task 1 result: %v\n", task1_result)
	fmt.Printf("Task 2 result: %v\n", task2_result)

}

func parseInput(scanner *bufio.Scanner) {

}

func task1() int {

	sum := 0

	return sum
}

func task2() int {

	sum := 0

	return sum
}
