package main

import (
	"regexp"
	"strconv"
)

// 183669043 59097164 (sample 161 48)
func main() {
	solve("03_input.txt", line03a, line03b, func() int { return total03a }, func() int { return total03b })
}

var (
	total03a   = 0
	total03b   = 0
	enabled03b = true
)

func line03a(line string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total03a += num1 * num2
	}
}

func line03b(line string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if match[0] == "do()" {
			enabled03b = true
			continue
		}
		if match[0] == "don't()" {
			enabled03b = false
			continue
		}
		if enabled03b {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total03b += num1 * num2
		}
	}
}
