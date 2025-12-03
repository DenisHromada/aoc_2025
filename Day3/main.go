package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltageSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		p_s, p_e := findBatteryPair(line)
		joltage := (p_s-'0')*10 + p_e - '0'

		joltageSum += int(joltage)
	}

	println(joltageSum)
}

func findBatteryPair(pack string) (byte, byte) {
	best_pair_start := pack[0]
	best_pair_end := pack[1]
	for i := 2; i < len(pack); i++ {

		if best_pair_end > best_pair_start {
			best_pair_start = best_pair_end
			best_pair_end = pack[i]
			continue
		}

		if pack[i] > best_pair_end {
			best_pair_end = pack[i]
		}
	}
	return best_pair_start, best_pair_end
}
