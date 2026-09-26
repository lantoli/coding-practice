package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 278 569808947758890 (sample 6 16)
func main() {
	solve19(false)
	solve19(true)
}

var (
	file19 = "19_input.txt"
)

func solve19(b bool) {
	f, _ := os.Open("inputs/" + file19)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	total := 0
	cache := map[string]int{"": 1}
	for scanner.Scan() {
		total += calc19(patterns, scanner.Text(), cache, b)
	}
	fmt.Println(total)
}

func calc19(patterns []string, line string, cache map[string]int, b bool) int {
	val, found := cache[line]
	if found {
		return val
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(line, pattern) {
			val += calc19(patterns, line[len(pattern):], cache, b)
		}
	}
	if !b && val > 1 {
		val = 1
	}
	cache[line] = val
	return val
}
