package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type All struct {
	source    image.Point
	splitters map[int]map[int]bool
	depth     int
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
	in := All{}
	in.splitters = make(map[int]map[int]bool)

	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			curPoint := image.Point{i, lineNumber}
			switch c {
			case 'S':
				in.source = curPoint
			case '^':
				if in.splitters[lineNumber] == nil {
					in.splitters[lineNumber] = make(map[int]bool)
				}
				in.splitters[lineNumber][i] = true
			}
		}

		lineNumber++
	}
	in.depth = lineNumber

	return in
}

func task1(parsedInput All) int {
	sum := 0

	down := image.Point{0, 1}

	initialBeam := parsedInput.source.Add(down)
	beams := map[int]bool{}
	beams[initialBeam.X] = true

	for i := range parsedInput.depth {
		thisRowSplitters, rowHasSplitters := parsedInput.splitters[i]
		for j, active := range beams {
			if !active {
				continue
			}
			if rowHasSplitters {
				splitter, exists := thisRowSplitters[j]
				if exists && splitter {
					beams[j] = false
					beams[j-1] = true
					beams[j+1] = true
					sum++
				}
			}
		}
	}

	return sum
}

func task2(parsedInput All) int {
	sum := 0

	down := image.Point{0, 1}

	initialBeam := parsedInput.source.Add(down)
	beams := map[int]int{}
	beams[initialBeam.X] = 1

	for i := range parsedInput.depth {
		thisRowSplitters, rowHasSplitters := parsedInput.splitters[i]
		for j, active := range beams {
			if active == 0 {
				continue
			}
			if rowHasSplitters {
				splitter, exists := thisRowSplitters[j]
				if exists && splitter {

					beams[j-1] += beams[j]
					beams[j+1] += beams[j]
					beams[j] = 0
				}
			}
		}
	}

	for _, value := range beams {
		sum += value
	}

	return sum
}
