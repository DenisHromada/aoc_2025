package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const circle_size = 100

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	current_position := 50
	current_position_sanity := 50
	zero_clicks := 0
	zero_clicks_sanity := 0

	for scanner.Scan() {
		if current_position < 0 || current_position >= circle_size {
			panic("Current position out of bounds")
		}

		line := scanner.Text()

		direction := line[0]
		value, _ := strconv.Atoi(line[1:])

		change := 1
		if direction == 'L' {
			change = -1
		}

		for i := 0; i < value; i++ {
			current_position_sanity += change
			if current_position_sanity == 0 || current_position_sanity == circle_size {
				current_position_sanity = 0
				zero_clicks_sanity++
			} else if current_position_sanity < 0 {
				current_position_sanity = circle_size - 1
			}
		}

		if value == 0 {
			continue
		}

		extra_rotations := value / circle_size
		normalized_value := value % circle_size

		zero_clicks += extra_rotations

		if direction == 'L' {
			normalized_value = -normalized_value
		}

		new_position := (current_position + normalized_value + circle_size) % circle_size

		if (current_position != 0) && ((new_position == 0) || (current_position+normalized_value < 0) || (current_position+normalized_value >= circle_size)) {
			zero_clicks++
		}
		current_position = new_position

	}

	fmt.Printf("mine %v; sanity %v", zero_clicks, zero_clicks_sanity)
}
