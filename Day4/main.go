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

func (w Warehouse) is_in_bounds(place Point) bool {
	if 0 <= place.Y && place.Y < w.height {
		if 0 <= place.X && place.X < w.width {
			return true
		}
	}
	return false
}

func (w Warehouse) getVal(place Point) bool {
	return w.floor[place.Y][place.X]
}

func main() {

	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	warehouseMap := parseWarehouse(scanner)
	fmt.Printf("Warehouse width: %v & height: %v\n", warehouseMap.width, warehouseMap.height)

	task1_result := task1(warehouseMap)

	fmt.Printf("Task 1 result: %v\n", task1_result)

}

func task1(warehouse Warehouse) int {

	condition_sum := 0

	for y := range warehouse.height {
		for x := range warehouse.width {
			place := Point{x, y}
			if warehouse.getVal(place) && is_not_surrounded(warehouse, place) {
				condition_sum++
			}
		}
	}

	return condition_sum
}

func is_not_surrounded(warehouse Warehouse, place Point) bool {
	neighbors := 0

	neighborhood_delta := [8]Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for k := range len(neighborhood_delta) {
		neighbor_place := place.Add(neighborhood_delta[k])
		if warehouse.is_in_bounds(neighbor_place) {
			if warehouse.getVal(neighbor_place) {
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

		warehouse_row := make([]bool, len(line))
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '@':
				warehouse_row[i] = true
			case '.':
				warehouse_row[i] = false
			default:
				panic("oops")
			}
		}
		warehouse = append(warehouse, warehouse_row)
	}

	w_struct := Warehouse{warehouse, len(warehouse[0]), len(warehouse)}

	return w_struct
}
