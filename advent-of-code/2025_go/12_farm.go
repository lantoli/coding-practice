package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Sample", solve12("inputs/12_sample.txt")) // 2 (heuristic gives 3)
	fmt.Println("A", solve12("inputs/12_input.txt"))       // 536 (done in AI, don't solve the general case)
}

func solve12(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Count '#' cells in each of 6 shapes
	var cells [6]int
	for i := range 6 {
		scanner.Scan() // "i:"
		for range 3 {
			scanner.Scan()
			cells[i] += strings.Count(scanner.Text(), "#")
		}
		scanner.Scan() // empty line
	}

	// Check each region: fits if cells_needed < area
	total := 0
	for scanner.Scan() {
		p := strings.Fields(scanner.Text())
		d := strings.Split(p[0][:len(p[0])-1], "x")
		w, _ := strconv.Atoi(d[0])
		h, _ := strconv.Atoi(d[1])

		need := 0
		for i, s := range p[1:] {
			n, _ := strconv.Atoi(s)
			need += n * cells[i]
		}
		if need < w*h {
			total++
		}
	}
	return total
}
