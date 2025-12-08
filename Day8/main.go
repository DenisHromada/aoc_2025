package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type All struct {
	points []Vector3
}

type Vector3 struct {
	X int
	Y int
	Z int
}

func (p1 Vector3) Subtract(p2 Vector3) Vector3 {
	return Vector3{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

func (p Vector3) Length() float64 {
	l := math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
	if l == 0 {
		panic("calculated zero distance")
	}
	return l
}

type Circuits struct {
	points map[Vector3]int
}

// func (c Circuit) Size() int {
// 	return len(c.points)
// }

func (c *Circuits) mergeByCircuitIds(circuitNumberA int, circuitNumberB int) bool {
	if circuitNumberA == circuitNumberB {
		return false
	}
	fmt.Printf("merging %v into %v\n", circuitNumberB, circuitNumberA)
	for p, circuitNumber := range c.points {
		if circuitNumber == circuitNumberB {
			c.points[p] = circuitNumberA
		}
	}
	return true
}

func (c *Circuits) mergeByPoints(pointA Vector3, pointB Vector3) bool {
	return c.mergeByCircuitIds(c.points[pointA], c.points[pointB])
}

type PointPair struct {
	points   [2]Vector3
	distance float64
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

	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()

		slices := strings.Split(line, ",")

		if len(slices) != 3 {
			panic("invalid point")
		}

		iX, errX := strconv.Atoi(slices[0])
		iY, errY := strconv.Atoi(slices[1])
		iZ, errZ := strconv.Atoi(slices[2])

		if errX != nil || errY != nil || errZ != nil {
			panic("could not parse to number")
		}

		newPoint := Vector3{iX, iY, iZ}
		in.points = append(in.points, newPoint)

		lineNumber++
	}

	return in
}

const task_1_steps = 1_000

// const task_1_steps = 10
const task_1_magic = 5_000
const task_1_top_circuits = 3

func task1(parsedInput All) int {
	product := 1

	circuits := Circuits{points: make(map[Vector3]int)}
	for i, p := range parsedInput.points {
		circuits.points[p] = i
	}
	// fmt.Printf("%v\n", circuits.points[parsedInput.points[500]])

	paired := 0
	for {
		closestPairs := []PointPair{}
		closestPairs = getNClosestPairs(parsedInput, circuits, closestPairs)
		sort.Slice(closestPairs, func(i int, j int) bool {
			return closestPairs[i].distance < closestPairs[j].distance
		})
		closestPairs = closestPairs[:min(len(closestPairs), task_1_steps)]

		// fmt.Printf("%v %v\n", testPoint, circuits.points[testPoint])
		// fmt.Printf("%v\n", closestPairs)

		// fmt.Printf("%v\n", circuits.points)
		for _, pair := range closestPairs {
			// fmt.Printf("point %v circuit %v\n", pair.points[0], circuits.points[pair.points[0]])
			if circuits.mergeByPoints(pair.points[0], pair.points[1]) {
				paired++
				if paired+1 == task_1_steps {
					break
				}
			}
		}
		if paired+1 == task_1_steps {
			break
		}

	}

	largest_circuits := map[int]int{}
	for _, circuitId := range circuits.points {
		largest_circuits[circuitId]++
	}

	topN := []int{}
	for _, value := range largest_circuits {
		topN = append(topN, value)
	}
	sort.Ints(topN)

	fmt.Printf("%v\n", topN)
	topN = topN[len(topN)-task_1_top_circuits:]
	fmt.Printf("%v\n", topN)

	for _, value := range topN {
		product *= value
	}

	return product
}

func getNClosestPairs(parsedInput All, circuits Circuits, closestPairs []PointPair) []PointPair {
	for i, pointA := range parsedInput.points {
		if i+1 >= len(parsedInput.points) {
			break
		}
		for _, pointB := range parsedInput.points[i+1:] {
			if circuits.points[pointA] == circuits.points[pointB] {
				continue
			}
			newPointPair := PointPair{points: [2]Vector3{pointA, pointB}, distance: pointA.Subtract(pointB).Length()}
			closestPairs = append(closestPairs, newPointPair)

			if len(closestPairs) >= task_1_steps*task_1_magic {
				sort.Slice(closestPairs, func(i int, j int) bool {
					return closestPairs[i].distance < closestPairs[j].distance
				})
				closestPairs = closestPairs[:task_1_steps]
			}
		}
	}
	return closestPairs
}

func task2(parsedInput All) int {
	sum := 0

	return sum
}
