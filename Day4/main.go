package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

type Warehouse struct {
	floor  [][]bool
	width  int
	height int //or depth/length
}

func (w Warehouse) IsInBounds(p Point) bool {
	if 0 <= p.Y && p.Y < w.height {
		if 0 <= p.X && p.X < w.width {
			return true
		}
	}
	return false
}

func (w Warehouse) Get(p Point) bool {
	return w.floor[p.Y][p.X]
}

func (w *Warehouse) Set(p Point, value bool) {
	w.floor[p.Y][p.X] = value
}

func main() {

	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	warehouseMap := parseWarehouse(scanner)
	fmt.Printf("Warehouse width: %v & height: %v\n", warehouseMap.width, warehouseMap.height)

	task1Result := task1(warehouseMap)
	task2Result := task2(warehouseMap)

	fmt.Printf("Task 1 result: %v\n", task1Result)
	fmt.Printf("Task 2 result: %v\n", task2Result)

}

func task1(warehouse Warehouse) int {

	conditionSum := 0

	for y := range warehouse.height {
		for x := range warehouse.width {
			place := Point{x, y}
			if warehouse.Get(place) && isNotSurrounded(warehouse, place) {
				conditionSum++
			}
		}
	}

	return conditionSum
}

func isNotSurrounded(warehouse Warehouse, place Point) bool {
	neighbors := 0

	neighborhoodDeltas := [8]Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for k := range len(neighborhoodDeltas) {
		neighborPlace := place.Add(neighborhoodDeltas[k])
		if warehouse.IsInBounds(neighborPlace) {
			if warehouse.Get(neighborPlace) {
				neighbors++
			}

		}
	}

	if neighbors < 4 {
		return true
	}
	return false

}

func parseWarehouse(scanner *bufio.Scanner) Warehouse {

	var warehouse [][]bool

	for scanner.Scan() {
		line := scanner.Text()

		warehouseRow := make([]bool, len(line))
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '@':
				warehouseRow[i] = true
			case '.':
				warehouseRow[i] = false
			default:
				panic("oops")
			}
		}
		warehouse = append(warehouse, warehouseRow)
	}

	w := Warehouse{warehouse, len(warehouse[0]), len(warehouse)}

	return w
}

func task2(warehouse Warehouse) int {

	accessed := 0
	prevAccessed := 0

	for {
		prevAccessed = accessed

		for y := range warehouse.height {
			for x := range warehouse.width {
				place := Point{x, y}
				if warehouse.Get(place) && isNotSurrounded(warehouse, place) {
					accessed++
					warehouse.Set(place, false)
				}
			}
		}
		if prevAccessed == accessed {
			break
		}
	}

	return accessed
}
