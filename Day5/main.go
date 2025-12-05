package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r Range) contains(id int) bool {
	return r.Start <= id && id <= r.End
}

type All struct {
	ranges        []Range
	ingredientIds []int
}

func main() {

	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsedInput := parseInput(scanner)

	task1_result := task1(parsedInput)
	task2_result := task2()

	fmt.Printf("Task 1 result: %v\n", task1_result)
	fmt.Printf("Task 2 result: %v\n", task2_result)

}

func parseInput(scanner *bufio.Scanner) All {
	input := All{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, "-")
		start, e1 := strconv.Atoi(nums[0])
		end, e2 := strconv.Atoi(nums[1])

		if e1 != nil || e2 != nil {
			panic("fialed to parse range")
		}

		new_range := Range{start, end}
		input.ranges = append(input.ranges, new_range)
	}

	for scanner.Scan() {
		line := scanner.Text()
		newId, e := strconv.Atoi(line)

		if e != nil {
			panic("failed to parse id")
		}

		input.ingredientIds = append(input.ingredientIds, newId)
	}

	return input
}

func task1(parsedInput All) int {
	sum := 0

	for i := 0; i < len(parsedInput.ingredientIds); i++ {
		for j := 0; j < len(parsedInput.ranges); j++ {
			if parsedInput.ranges[j].contains(parsedInput.ingredientIds[i]) {
				sum++
				break
			}
		}
	}

	return sum
}

func task2() int {

	sum := 0

	return sum
}
