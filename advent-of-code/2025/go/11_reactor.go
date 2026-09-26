package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Sample A", solve11("inputs/11_sample_a.txt", false)) // 5
	fmt.Println("A", solve11("inputs/11_input.txt", false))           // 599
	fmt.Println("Sample B", solve11("inputs/11_sample_b.txt", true))  // 2
	fmt.Println("B", solve11("inputs/11_input.txt", true))            // 393474305030400
}

type state struct {
	node string
	fft  bool
	dac  bool
}

func solve11(filename string, b bool) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		nodes   = make(map[string][]string)
		cache   = make(map[state]int)
	)
	defer f.Close()
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		key := parts[0][:len(parts[0])-1]
		for i := 1; i < len(parts); i++ {
			nodes[key] = append(nodes[key], parts[i])
		}
	}
	start := "you"
	if b {
		start = "svr"
	}
	return dfs(start, false, false, b, nodes, cache)
}

func dfs(node string, fft, dac, b bool, nodes map[string][]string, cache map[state]int) int {
	if node == "fft" {
		fft = true
	}
	if node == "dac" {
		dac = true
	}
	s := state{node, fft, dac}
	if v, ok := cache[s]; ok {
		return v
	}
	total := 0
	for _, next := range nodes[node] {
		if next == "out" {
			if !b || (fft && dac) {
				total++
			}
		} else {
			total += dfs(next, fft, dac, b, nodes, cache)
		}
	}
	cache[s] = total
	return total
}
