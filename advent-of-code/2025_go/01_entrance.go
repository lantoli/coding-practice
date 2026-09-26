package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sample A", solve01("inputs/01_sample.txt", false)) // 3
	fmt.Println("A", solve01("inputs/01_input.txt", false))         // 1084
	fmt.Println("Sample B", solve01("inputs/01_sample.txt", true))  // 6
	fmt.Println("B", solve01("inputs/01_input.txt", true))          // 6475
}

func solve01(filename string, b bool) int {
	var (
		f, _                     = os.Open(filename)
		scanner                  = bufio.NewScanner(f)
		dir                      rune
		total, num, mod, newDial int
		dial                     = 50
	)
	defer f.Close()
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%c%d", &dir, &num)
		mod = num % 100
		if dir == 'L' {
			mod = -mod
		}
		newDial = (dial + mod + 100) % 100
		if b {
			total += num / 100
			if dial != 0 && newDial != 0 && dial+mod != newDial {
				total++
			}
		}
		dial = newDial
		if dial == 0 {
			total++
		}
	}
	return total
}

// Initial solutions, not needed anymore

func solve01a(filename string) int {
	data, _ := os.ReadFile(filename)
	dial := 50
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) < 2 {
			break
		}
		left, numStr := line[0] == 'L', line[1:]
		num, _ := strconv.Atoi(numStr)
		if left {
			num = -num
		}
		dial = (dial + num) % 100
		if dial == 0 {
			total++
		}
	}
	return total
}

func solve01bbrute(filename string) int {
	data, _ := os.ReadFile(filename)
	dial := 50
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) < 2 {
			break
		}
		left, numStr := line[0] == 'L', line[1:]
		num, _ := strconv.Atoi(numStr)
		dir := 1
		if left {
			dir = -1
		}
		for range num {
			dial += dir
			if dial == -1 {
				dial = 99
			}
			if dial == 100 {
				dial = 0
			}
			if dial == 0 {
				total++
			}
		}
	}
	return total
}
