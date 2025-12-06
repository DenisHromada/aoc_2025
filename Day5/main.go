package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(id int) bool {
	return r.Start <= id && id <= r.End
}

func (r1 Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r1.End >= r2.Start
}

func (r1 Range) Merge(r2 Range) Range {
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

	task1Result := task1(parsedInput)
	task2Result := task2(parsedInput)

	fmt.Printf("Task 1 result: %v\n", task1Result)
	fmt.Printf("Task 2 result: %v\n", task2Result)

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

		newRange := Range{start, end}
		input.ranges = append(input.ranges, newRange)
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
			if parsedInput.ranges[j].Contains(parsedInput.ingredientIds[i]) {
				sum++
				break
			}
		}
	}

	return sum
}

func task2(parsedInput All) int {
	sum := 0

	mergedRanges := mergeRanges(parsedInput.ranges)

	for i := 0; i < len(mergedRanges); i++ {
		start := mergedRanges[i].Start
		end := mergedRanges[i].End
		// elements := end - start + 1
		// fmt.Printf("%v - %v is %v elements\n", start, end, elements)
		sum += end - start + 1
	}

	return sum
}

func mergeRanges(originalRanges []Range) []Range {
	sort.Slice(originalRanges, func(i int, j int) bool {
		return originalRanges[i].Start < originalRanges[j].Start
	})

	newRanges := []Range{originalRanges[0]}

	for i := 1; i < len(originalRanges); i++ {
		lastRange := &newRanges[len(newRanges)-1]
		if originalRanges[i].Start <= lastRange.End {
			lastRange.End = max(lastRange.End, originalRanges[i].End)
		} else {
			newRanges = append(newRanges, originalRanges[i])
		}
	}

	return newRanges
}
