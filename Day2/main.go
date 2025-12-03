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

	scanner.Split(SplitByComma)

	silly_accum := 0

	for scanner.Scan() {
		token := scanner.Text()

		token = strings.TrimSpace(token)
		slices := strings.Split(token, "-")

		start := slices[0]
		end := slices[1]
		if len(slices) != 2 {
			panic("Invalid token format: " + token)
		}

		start_num, _ := strconv.Atoi(start)
		end_num, _ := strconv.Atoi(end)

		if start_num > end_num {
			fmt.Printf("%v - %v\n", start_num, end_num)
			panic("Invalid range: " + token)
		}

		for id := start_num; id <= end_num; id++ {
			if isSilly_v2(id) {
				silly_accum += id
			}
		}
	}

	fmt.Printf("Silly sum: %v\n", silly_accum)
}

// func isSilly(id int) bool {
// 	num_str := strconv.Itoa(id)
// 	if len(num_str)%2 != 0 {
// 		return false
// 	}
// 	mid := len(num_str) / 2
// 	first_half := num_str[:mid]
// 	second_half := num_str[mid:]

// 	if first_half == second_half {
// 		return true
// 	}
// 	return false
// }

func isSilly_v2(id int) bool {
	num_str := strconv.Itoa(id)

	for i := 1; i <= len(num_str)/2; i++ {
		if has_silly_substring(num_str, i) {
			return true
		}
	}
	return false
}

func has_silly_substring(s string, subs_len int) bool {
	if len(s)%subs_len != 0 {
		return false
	}

	num_subs := len(s) / subs_len
	first_subs := s[0:subs_len]

	for i := 1; i < num_subs; i++ {
		if first_subs != s[i*subs_len:(i+1)*subs_len] {
			return false
		}
	}

	return true
}

func SplitByComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range data {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
