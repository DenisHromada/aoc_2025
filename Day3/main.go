package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltagePairSum := 0
	joltageDozenSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		p_s, p_e := findBatteryPair(line)
		joltage := (p_s-'0')*10 + p_e - '0'

		joltagePairSum += int(joltage)

		pack_dozen := findBaterryDozen(line)
		dozen_joltage, err := strconv.Atoi(pack_dozen)
		if err != nil {
			panic(err)
		}
		joltageDozenSum += dozen_joltage

	}

	println("pair sum: ", joltagePairSum)
	println("dozen sum: ", joltageDozenSum)
}

func findBaterryDozen(pack string) string {
	pack_len := len(pack)
	best_dozen := []byte(pack[pack_len-12:])

	for i := pack_len - 13; i >= 0; i-- {

		candidate := pack[i]
		fmt.Printf("got pack: %s; and candidate: %s\n", string(best_dozen), string(candidate))

		for j := 0; j < 12; j++ {
			if candidate >= best_dozen[j] {
				candidate, best_dozen[j] = best_dozen[j], candidate
			} else {
				break
			}
		}
		fmt.Printf("MODIFIED and got pack: %s\n", string(best_dozen))

	}

	return string(best_dozen)
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
