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

type All2 struct {
	lines     [][]byte
	operators []byte
}

func main() {

	// file, _ := os.Open("input_sample")
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	file2, _ := os.Open("input")
	defer file2.Close()
	scanner2 := bufio.NewScanner(file2)

	parsedInput := parseInput(scanner)
	parsedInput2 := parseInput2(scanner2)

	// fmt.Printf(string(parsedInput.operators))

	task1Result := task1(parsedInput)
	task2Result := task2(parsedInput2)

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

	// for _, line := range input.lines {
	// 	for _, el := range line {
	// 		fmt.Printf("%v ", el)
	// 	}
	// 	fmt.Printf("\n")
	// }
	// for _, el := range input.operators {
	// 	fmt.Printf("%v ", el)
	// }
	// fmt.Printf("\n")

	return input
}

func parseInput2(scanner *bufio.Scanner) All2 {
	input := All2{}

	for scanner.Scan() {
		line := scanner.Text()
		input.lines = append(input.lines, []byte(line))
	}

	opRow := input.lines[len(input.lines)-1]
	operators := strings.Fields(string(opRow))
	for _, op := range operators {
		input.operators = append(input.operators, op[0])
	}

	input.lines = input.lines[:len(input.lines)-1]

	// fmt.Printf("%v %v; %v\n", len(input.lines), len(input.lines[0]), len(input.operators))

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

	// fmt.Printf("%v\n", results)
	for _, num := range results {
		sum += num
	}

	return sum
}

func task2(parsedInput All2) int {
	sum := 0

	curOpIndex := 0
	isFirstForThisOp := true
	results := make([]int, len(parsedInput.lines[0]))

	for x := range len(parsedInput.lines[0]) {
		curNum := 0
		colIsEmpty := true
		for y := range len(parsedInput.lines) {
			ch := parsedInput.lines[y][x]
			if ch == ' ' {
				continue
			}
			colIsEmpty = false

			n := int(ch - '0')
			curNum = curNum*10 + n
			// fmt.Printf("%v\n", curNum)
		}

		if isFirstForThisOp {
			results[curOpIndex] = curNum
			isFirstForThisOp = false
		} else if !colIsEmpty {
			switch parsedInput.operators[curOpIndex] {
			case '*':
				// fmt.Printf("%v * %v\n", results[curOpIndex], curNum)
				results[curOpIndex] *= curNum
			case '+':
				// fmt.Printf("%v + %v\n", results[curOpIndex], curNum)
				results[curOpIndex] += curNum
			}
		}

		if colIsEmpty {
			// fmt.Printf("break\n")
			curOpIndex++
			isFirstForThisOp = true
		}

	}

	for _, num := range results {
		sum += num
	}

	return sum
}
