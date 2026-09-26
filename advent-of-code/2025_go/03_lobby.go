package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sample A", solve03("inputs/03_sample.txt", 2))  // 357
	fmt.Println("A", solve03("inputs/03_input.txt", 2))          // 17095
	fmt.Println("Sample B", solve03("inputs/03_sample.txt", 12)) // 3121910778619
	fmt.Println("B", solve03("inputs/03_input.txt", 12))         // 168794698570517
}

func solve03(filename string, size int) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		total   = 0
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		total += solve03Line(line, size)
	}
	return total
}

func solve03Line(line string, size int) int {
	result := 0
	for i := 0; i < size; i++ {
		id := 0
		for j := 1; j < len(line)-(size-i-1); j++ {
			if line[j] > line[id] {
				id = j
			}
		}
		result = result*10 + int(line[id]-'0')
		line = line[id+1:]
	}
	return result
}
