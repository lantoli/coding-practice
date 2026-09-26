package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sample A", solve07a("inputs/07_sample.txt", false)) // 21
	fmt.Println("A", solve07a("inputs/07_input.txt", false))         // 1678
	fmt.Println("Sample B", solve07a("inputs/07_sample.txt", true))  // 40
	fmt.Println("B", solve07a("inputs/07_input.txt", true))          // 357525737893560
}

func solve07a(filename string, b bool) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		splits  = 0
		count   []int
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if count == nil {
			count = make([]int, len(line))
		}
		for x := range line {
			switch line[x] {
			case 'S':
				count[x] = 1
			case '^':
				count[x-1] += count[x]
				count[x+1] += count[x]
				if count[x] > 0 {
					splits++
				}
				count[x] = 0
			}
		}
	}
	if b {
		paths := 0
		for x := range count {
			paths += count[x]
		}
		return paths
	}
	return splits
}
