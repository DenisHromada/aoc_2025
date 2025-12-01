package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const circle_size = 100

func main() {
	inputFile, _ := os.Open("input")
	defer inputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)

	pos := 50
	pos_0_count := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()

		direction := line[0]
		change, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'L':
			pos = (pos - change) % circle_size
		case 'R':
			pos = (pos + change) % circle_size
		}

		if pos == 0 {
			pos_0_count++
		}
	}

	fmt.Println(pos_0_count)
}
