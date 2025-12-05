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

func (r1 Range) overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r1.End >= r2.Start
}

func (r1 Range) merge(r2 Range) Range {
	return Range{min(r1.Start, r2.Start), max(r1.End, r2.End)}
}

type All struct {
	ranges        []Range
	ingredientIds []int
}

func main() {

	// file, _ := os.Open("input_sample")
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsedInput := parseInput(scanner)

	task1_result := task1(parsedInput)
	task2_result := task2(parsedInput)

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

func task2(parsedInput All) int {
	sum := 0

	merged_ranges, hasMerged := mergeRangesIteration(parsedInput.ranges)
	for {
		if hasMerged == false {
			break
		}
		merged_ranges, hasMerged = mergeRangesIteration(merged_ranges)
	}

	for i := 0; i < len(merged_ranges); i++ {
		start := merged_ranges[i].Start
		end := merged_ranges[i].End
		elements := end - start + 1
		fmt.Printf("%v - %v is %v elements\n", start, end, elements)
		sum += end - start + 1
	}

	return sum
}

func mergeRangesIteration(original_ranges []Range) ([]Range, bool) {
	new_ranges := []Range{}
	has_merged := false
	for i := 0; i < len(original_ranges); i++ {
		originalRange := original_ranges[i]
		merged := false
		for j := 0; j < len(new_ranges); j++ {
			newR := new_ranges[j]
			if newR.overlaps(originalRange) {
				fmt.Printf("%v-%v & %v-%v", originalRange.Start, originalRange.End, newR.Start, newR.End)
				new_ranges[j] = newR.merge(originalRange)
				fmt.Printf(" => %v-%v\n", new_ranges[j].Start, new_ranges[j].End)
				merged = true
				has_merged = true
			}
		}
		if merged == false {
			new_ranges = append(new_ranges, originalRange)
		}
	}
	return new_ranges, has_merged
}
